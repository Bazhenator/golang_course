package models

import (
	"os"
	"strings"

	u "task_7/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token" ;sql:"-"`
}

func (account *Account) ValidateAccount() (map[string]interface{}, bool) {

	if !IsValidEmail(account.Email) {
		return u.Message(false, "Invalid email address"), false
	}

	if !IsEmailUnique(account.Email) {
		return u.Message(false, "Email is already in use"), false
	}

	if !IsValidPassword(account.Password) {
		return u.Message(false, "Password must be at least 8 characters long and contain at least one digit and one uppercase letter"), false
	}

	return u.Message(false, "Check is passed!"), true
}

func (account *Account) CreateAccount() map[string]interface{} {
	if resp, ok := account.ValidateAccount(); !ok {
		return resp
	}

	//pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	//account.Password = string(pwd)
	stars := strings.Repeat("*", len(account.Password))
	account.Password = stars

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	GetDB().Model(&account).Update("token", account.Token)

	response := u.Message(true, "Account has been created!")
	response["token"] = account.Token
	response["account"] = account
	return response
}

func LoginAccount(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password does not match!")
	}
	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func (account *Account) UpdateAccount() map[string]interface{} {
	if account.ID == 0 {
		return u.Message(false, "Account ID is required for update")
	}

	if response, ok := account.ValidateAccount(); !ok {
		return response
	}

	existingAccount := GetUser(account.ID)
	if existingAccount == nil {
		return u.Message(false, "Account not found")
	}

	existingAccount.Email = account.Email

	GetDB().Save(existingAccount)

	resp := u.Message(true, "Account has been updated successfully")
	resp["account"] = existingAccount
	return resp
}

func DeleteAccount(id uint) map[string]interface{} {
	account := GetUser(id)
	if account == nil {
		return u.Message(false, "Account not found")
	}

	GetDB().Where("user_id = ?", id).Delete(&Contact{})

	GetDB().Delete(account)

	resp := u.Message(true, "Account and associated contacts have been deleted successfully")
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}

	return acc
}

func GetAllUsers() []*Account {
	users := make([]*Account, 0)
	GetDB().Table("accounts").Find(&users)

	for _, user := range users {

		var temp struct {
			Token string
		}
		GetDB().Table("accounts").Where("id = ?", user.ID).Select("token").Scan(&temp)
		user.Token = temp.Token
	}

	return users
}
