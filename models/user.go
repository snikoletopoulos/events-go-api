package models

import "example.com/events-rest-api/db"

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

	result, err := statement.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	return err
}
