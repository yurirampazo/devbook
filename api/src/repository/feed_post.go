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
// Gets a post by its id
func (repo FeedPost) GetById(id uint64) (model.FeedPost, error) {
	line, err := repo.db.Query(`SELECT p.*, u.nick FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?`, id)
	if err != nil {
		return model.FeedPost{}, err
	}
	defer line.Close()

	var post model.FeedPost

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return model.FeedPost{}, err
		}
	}
	return post, nil
}
