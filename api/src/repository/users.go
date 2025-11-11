package repository

import (
	"api/src/model"
	"database/sql"
)

// user repository
type Users struct {
	db *sql.DB
}

// create user repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create an user
func (u Users) Create(user model.User) (uint64, error) {
	return 0, nil
}