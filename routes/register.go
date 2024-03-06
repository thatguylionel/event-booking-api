package routes

import (
	"net/http"
	"strconv"
	"tgl/eventapi/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err})
		return
	}

	err = event.CancelRegistration(userID, eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Your registration for the event has been cancelled successfully"})
}
