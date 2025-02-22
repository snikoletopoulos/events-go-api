package models

import (
	"errors"

	"events-rest-api/db"
	"events-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	statement, err := db.DB.Prepare(`
    INSERT INTO users (email, password) VALUES (?, ?)
    `)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	return err
}

func (user *User) ValidateCredentials() error {
	row := db.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", user.Email)

	var userPassword string
	if err := row.Scan(&user.ID, &userPassword); err != nil {
		return errors.New("invalid credentials")
	}

	isPasswordValid := utils.ComparePasswordHash(user.Password, userPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	user.Password = userPassword

	return nil
}
