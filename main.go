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

	r.HandleFunc("/create-user", service.CreateUser).Methods("POST")
	r.HandleFunc("/get-user", service.GetUser).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
