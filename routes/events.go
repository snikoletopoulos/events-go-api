package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get events"})
		return
	}
	if events == nil {
		events = []models.Event{}
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := models.FindEventByID(uint(eventID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

type EventBody struct {
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func createEvent(context *gin.Context) {
	var eventData EventBody
	if err := context.ShouldBindJSON(&eventData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := context.GetUint("userID")
	fmt.Println("🪚 userID route:", userID)

	event := models.Event{
		Name:        eventData.Name,
		Description: eventData.Description,
		Location:    eventData.Location,
		DateTime:    eventData.DateTime,
		UserID:      uint(userID),
	}

	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

// TODO: make fields optional
func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := models.FindEventByID(uint(eventID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
		return
	}

	userID := context.GetUint("userID")
	fmt.Println("🪚 userID:", userID)
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	var eventData EventBody
	if err := context.ShouldBindJSON(&eventData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event data"})
		return
	}

	event.Name = eventData.Name
	event.Description = eventData.Description
	event.Location = eventData.Location
	event.DateTime = eventData.DateTime

	if err := event.Update(); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event id"})
		return
	}

	event, err := models.FindEventByID(uint(eventID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
		return
	}

	userID := context.GetUint("userID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	if err := event.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
