package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"tgl/eventapi/models"

	"github.com/gin-gonic/gin"
)

func PublicRouter() {}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting events", "error": err})
	}
	fmt.Println(events)
	context.JSON(http.StatusOK, events)

}

func getEventByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating event", "error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating event", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	userId := context.GetInt64("userId")
	deleteEvent, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err})
		return
	}

	if deleteEvent.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete this event"})
		return
	}

	err = deleteEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting event", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
