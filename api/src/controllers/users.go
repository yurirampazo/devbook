package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUsersRepository(db)
	repository.Create(user)
}

// get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users"))
}

// get user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting User"))
}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
