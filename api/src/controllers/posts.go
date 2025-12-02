package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
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
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewFeedPostsRepository(db)
	posts, err := repo.FindAllPosts(userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
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

// Update some post
func EditPost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}
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

	savedPost, err := repo.GetById(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPost.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("Update another user post is impossible, you shall not do it!"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.FeedPost

	if err = json.Unmarshal(reqBody, &post); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = repo.UpdatePost(postID, post); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Delete some post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}
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

	savedPost, err := repo.GetById(postID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPost.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("Delete another user post is impossible, you shall not do it!"))
		return
	}

	if err = repo.DeletePost(postID); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Finds all posts from an specific user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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

	posts, err := repo.GetUserPosts(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}
