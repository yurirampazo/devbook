package model

import (
	"errors"
	"strings"
	"time"
)

// User represents user utilizing a social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

//Prepare will validate and format user received data
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("Name cannot be empty")
	}
	
	if user.Nick == "" {
		return errors.New("Nick cannot be empty")
	}	
	
	if user.Password == "" {
		return errors.New("Password cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}