package utils

import (
	"net/http"
)

// Error structure for storing error information
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(e Error) map[string]interface{} {
	data := map[string]interface{}{
		"error": e,
	}
	return data
}

func BadRequest(w http.ResponseWriter) {
	e := Error{
		HTTPCode: http.StatusBadRequest,
		Code:     400,
		Message:  "Invalid request!",
	}
	w.WriteHeader(e.HTTPCode)
	Respond(w, JSONError(e))
}

func ServerError(w http.ResponseWriter) {
	e := Error{
		HTTPCode: http.StatusInternalServerError,
		Code:     500,
		Message:  "The server has encountered a situation it does not know how to handle :(",
	}
	w.WriteHeader(e.HTTPCode)
	Respond(w, JSONError(e))
}

func AuthorizationError(w http.ResponseWriter, message string) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     403,
		Message:  message,
	}
	w.WriteHeader(e.HTTPCode)
	Respond(w, JSONError(e))
}
