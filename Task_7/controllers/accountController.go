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
		u.BadRequest(w)
		return
	}
	resp := account.CreateAccount()
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}

var LoginAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.BadRequest(w)
		return
	}
	resp := models.LoginAccount(account.Email, account.Password)
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}

var UpdateAccount = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.BadRequest(w)
		return
	}

	if account.ID != userID {
		u.BadRequest(w)
		return
	}

	resp := account.UpdateAccount()
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	resp := models.DeleteAccount(userID)
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}

var GetUserByID = func(w http.ResponseWriter, r *http.Request) {
	user := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.BadRequest(w)
		return
	}

	existingUser := models.GetUser(user.ID)
	if existingUser != nil {
		resp := u.Message(true, "user is already exists")
		u.Respond(w, resp)
		return
	}

	u.BadRequest(w)
}

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	resp := u.Message(true, "success")
	resp["users"] = users
	u.Respond(w, resp)
}
