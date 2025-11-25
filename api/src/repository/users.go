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

// Get users by Name or nick
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

//Finds User BY Id
func (repository *Users) FindByID(userId uint64) (model.User, error) {
	lines, err := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", userId)
	if err != nil {
		return model.User{}, err
	}
	defer lines.Close()

	var user model.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}

// Update user
func (repository *Users) UpdateUser(user model.User, userID uint64) error {
	statement , err := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// Delete user by id
func (repository *Users) DeleteById(userID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?",)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// Finds registered user by email
func (repository *Users) FindByEmail(email string) (model.User, error) {
	line, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return model.User{}, err
	}
	defer line.Close()

	var user model.User

	if line.Next() {
		if err = line.Scan(&user.ID,&user.Password); err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}

//Follow another user
func (userRepository *Users) Follow (userID, followerID uint64) error {
	statement, err := userRepository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow someone, is that simple.
func (userRepository *Users) Unfollow (userID, followerID uint64) error {
	statement, err := userRepository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
		if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

// Find followers
func (userRepository *Users) FindFollowers (userID uint64) ([]model.User, error) {
	lines, err := userRepository.db.Query(`
	
	SELECT u.id, u.name, u.nick, u.email, u.createdAt
	FROM users u INNER JOIN followers f ON
	u.id = f.follower_id WHERE f.user_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []model.User
	

	for lines.Next() {
		var follower model.User

		if err = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)

	}	
	return followers, nil
}