package service

import (
	"demo/auth"
	"demo/db"
	"demo/resources"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input resources.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	token, err := auth.GenerateJwtToken(input.EmailId)
	if err != nil {
		fmt.Println("error while generating token")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("error while generating hash password")
	}
	input.Password = string(hash)
	err = db.CreateUser(input)
	if err != nil {
		fmt.Println("err in db")
	}
	fmt.Println(token)
	w.WriteHeader(201)
	w.Header().Add("token", token)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("token") == "" {
		w.WriteHeader(403)
	}
	var input resources.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	userdetails, err := db.GetUser(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	if CheckPasswordHash(userdetails.Password, input.Password) {
		userdetails.Password = "***********"
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(userdetails)
		return
	} else {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode("email or password wrong")
		return
	}

}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
