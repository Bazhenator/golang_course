package controllers

import (
	"encoding/json"
	"net/http"

	"strconv"
	"task_7/models"
	u "task_7/utils"

	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.ServerError(w)
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.BadRequest(w)
		return
	}

	contact.UserId = user
	resp := contact.UpdateContact()
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		u.BadRequest(w)
		return
	}

	resp := models.DeleteContact(uint(contactID))
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("user").(uint)
	if !ok {
		u.BadRequest(w)
		return
	}

	if models.GetUser(id) == nil {
		u.BadRequest(w)
		return
	}

	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
