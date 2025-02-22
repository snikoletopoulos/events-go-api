package routes

import (
	"net/http"

	"example.com/events-rest-api/models"
	"example.com/events-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body."})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body."})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials."})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
