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
	statement, err := 
	u.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?,?,?,?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)	
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil

}