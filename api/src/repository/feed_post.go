package repository

import (
	"api/src/model"
	"database/sql"
)

type FeedPost struct {
	db *sql.DB
}

// Creates a new  FeedPost Repository
func NewFeedPostsRepository(db *sql.DB) *FeedPost {
	return &FeedPost{db}
}

// Inserts a new post into database
func (repo FeedPost) Create(post model.FeedPost) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertedID), nil
}
