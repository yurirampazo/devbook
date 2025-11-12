package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

//Get users by Name or nick
func (repository Users) Find(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // real format %nameOrNick%

	lines, err := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", 
	nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []model.User

	for lines.Next() {
		var user model.User
		
		if err = lines.Scan(
			&user.ID, 
			&user.Name, 
			&user.Nick, 
			&user.Email, 
			&user.CreatedAt, 
		); err != nil {
			return nil, err
		} 

		users = append(users, user)
	}
	return users, nil
}