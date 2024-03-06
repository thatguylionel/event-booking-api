package main

import (
	"fmt"
	"net/http"
	"tgl/eventapi/db"
	"tgl/eventapi/models"

	"github.com/gin-gonic/gin"
)

const serverAddress = "localhost:8080"
const basePath = ""

func main() {
	db.InitDB()
	server := gin.Default()
	eventsRouter(server)

	server.Run(serverAddress)
}

const eventsPath = basePath + "/events"

func eventsRouter(server *gin.Engine) {
	server.GET(eventsPath, getEvents)
	server.POST(eventsPath, createEvent)
}

func PublicRouter() {}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting events", "error": err})
	}
	fmt.Println(events)
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating event", "error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
