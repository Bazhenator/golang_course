package utils

import (
	"encoding/json"
	"net/http"
)

// Error structure for storing error information
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// JSONError function to handle JSON error response
func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	_, _ = w.Write(b)
}

// NewError creates a new Error instance
func NewError(httpCode int, code int, message string) Error {
	return Error{
		HTTPCode: httpCode,
		Code:     code,
		Message:  message,
	}
}
