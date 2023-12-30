package main

import (
	"fmt"
	"net/http"
	"os"
	app "task_7/authentication"
	"task_7/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginAccount).Methods("POST")
	router.HandleFunc("/user/check", controllers.GetUserByID).Methods("POST")
	router.HandleFunc("/user/update", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/user/delete/{id:[0-9]+}", controllers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/update", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/delete/{id:[0-9]+}", controllers.DeleteContact).Methods("DELETE")
	router.HandleFunc("/me/contacts", controllers.GetContacts).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
