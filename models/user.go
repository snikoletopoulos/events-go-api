package models

import (
	"errors"

	"events-rest-api/db"
	"events-rest-api/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func FindByEmail(email string) (*User, error) {
	var user User
	result := db.DB.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindByID(id string) (*User, error) {
	var user User
	result := db.DB.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (user *User) Save() error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	result := db.DB.Create(&user)
	return result.Error
}

func (user *User) ValidateCredentials(inputPassword string) error {
	isPasswordValid := utils.ComparePasswordHash(user.Password, inputPassword)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
