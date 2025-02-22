package routes

import (
	"net/http"
	"strconv"

	"events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful."})
}

func cancelRegistrationForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancel registration successful."})
}
