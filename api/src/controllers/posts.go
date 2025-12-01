package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Ads a new post in database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var feedPost model.FeedPost
	if err = json.Unmarshal(reqBody, &feedPost); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	feedPost.AuthorID = userID

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewFeedPostsRepository(db)
	feedPost.ID, err = repo.Create(feedPost)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, feedPost)

}

// Gets the user feed user posts
func FindPosts(w http.ResponseWriter, r *http.Request) {

}

// Finds an especific post
func GetPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewFeedPostsRepository(db)
	post, err := repo.GetById(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

// UPdate some post
func EditPost(w http.ResponseWriter, r *http.Request) {

}

// Delete some post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
