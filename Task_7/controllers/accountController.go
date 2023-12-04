package controllers

import (
	"encoding/json"
	"net/http"
	"task_7/models"
	u "task_7/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		resp := u.NewError(http.StatusBadRequest, 400, "Invalid request")
		u.JSONError(w, resp)
		return
	}
	resp := account.CreateAccount()
	u.Respond(w, resp)
}

var LoginAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		resp := u.NewError(http.StatusBadRequest, 400, "Invalid request")
		u.JSONError(w, resp)
		return
	}
	resp := models.LoginAccount(account.Email, account.Password)
	u.Respond(w, resp)
}

var UpdateAccount = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		resp := u.NewError(http.StatusBadRequest, 400, "Invalid request")
		u.JSONError(w, resp)
		return
	}

	if account.ID != userID {
		resp := u.NewError(http.StatusBadRequest, 400, "You can only update your own account")
		u.JSONError(w, resp)
		return
	}

	resp := account.UpdateAccount()
	u.Respond(w, resp)
}

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	resp := models.DeleteAccount(userID)
	u.Respond(w, resp)
}

var GetUserByID = func(w http.ResponseWriter, r *http.Request) {
	user := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		resp := u.NewError(http.StatusBadRequest, 400, "Invalid request")
		u.JSONError(w, resp)
		return
	}

	existingUser := models.GetUser(user.ID)
	if existingUser != nil {
		resp := u.Message(true, "User exists")
		u.Respond(w, resp)
		return
	}

	resp := u.NewError(http.StatusBadRequest, 400, "User does not exist")
	u.JSONError(w, resp)
}

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	resp := u.Message(true, "success")
	resp["users"] = users
	u.Respond(w, resp)
}
