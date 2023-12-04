package models

import (
	"fmt"
	u "task_7/utils"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {

	if !IsValidName(contact.Name) {
		return u.Message(false, "Invalid contact name"), false
	}

	if !IsValidPhoneNumber(contact.Phone) {
		return u.Message(false, "Invalid phone number"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User not found!"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func (contact *Contact) UpdateContact() map[string]interface{} {
	if contact.ID == 0 {
		return u.Message(false, "Contact ID is required for update")
	}

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	existingContact := GetContact(contact.ID)
	if existingContact == nil {
		return u.Message(false, "Contact not found")
	}

	existingContact.Name = contact.Name
	existingContact.Phone = contact.Phone

	GetDB().Save(existingContact)

	resp := u.Message(true, "Contact has been updated successfully")
	resp["contact"] = existingContact
	return resp
}

func DeleteContact(id uint) map[string]interface{} {
	contact := GetContact(id)
	if contact == nil {
		return u.Message(false, "Contact not found")
	}

	GetDB().Delete(contact)

	resp := u.Message(true, "Contact has been deleted successfully")
	return resp
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}
