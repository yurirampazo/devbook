package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// AUthenticates user into API
func Login(w http.ResponseWriter, r *http.Request) { 
	reqBody, err := io.ReadAll(r.Body) 
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	savedUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	} 

	if err = security.VerifyPassword(savedUser.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.GenerateToken(savedUser.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}