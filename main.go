package main

import (
	"demo/db"
	"demo/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.DbConnect()

	r := mux.NewRouter()
	r.HandleFunc("/register", service.CreateUser).Methods("POST")
	r.HandleFunc("/login", service.Login).Methods("POST")

	// Admin routes
	adminRouter := r.PathPrefix("/admin").Subrouter()

	adminRouter.HandleFunc("/users", service.AdminListUsers).Methods("GET")
	adminRouter.HandleFunc("/users/{username}", service.AdminSearchUser).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
