package routes

import (
	"net/http"

	"events-rest-api/models"
	"events-rest-api/utils"

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
	var userData models.User
	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body."})
		return
	}

	user, err := models.FindByEmail(userData.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials."})
		return
	}

	isPasswordValid := utils.ComparePasswordHash(user.Password, userData.Password)
	if !isPasswordValid {
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
