package service

import (
	"demo/db"

	"demo/middleware"
	"demo/resources"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input resources.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	token, err := middleware.GenerateJwtToken(input.EmailId)
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password hash:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}
	input.Password = string(hash)

	err = db.CreateUser(input)
	if err != nil {
		fmt.Println("Error creating user in the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	fmt.Println("Token:", token)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("token", token)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	err := middleware.ExtractUserFromToken(r)
	if err != nil {
		fmt.Println("Error extracting user from token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var input resources.User
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}

	userdetails, err := db.GetUser(input)
	if err != nil {
		fmt.Println("Error getting user from the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	if CheckPasswordHash(userdetails.Password, input.Password) {
		userdetails.Password = "***********"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userdetails)
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Email or password is incorrect")
	}
}

// Admin: Display list of users
func AdminListUsers(w http.ResponseWriter, r *http.Request) {
	userdetails, err := db.GetUsers()
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	users := make([]resources.User, 0, len(userdetails))
	for _, u := range userdetails {
		users = append(users, u)
	}

	fmt.Println(users, "userdetails")
	json.NewEncoder(w).Encode(users)
}

// Admin: Search user by username
func AdminSearchUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userdetails, err := db.GetUserByName(username)
	if err != nil {
		fmt.Println("Error retrieving user by username:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	fmt.Println(userdetails, "userdetails")
	if userdetails.UserId == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	json.NewEncoder(w).Encode(userdetails)
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// if err !=
	return err == nil
}
