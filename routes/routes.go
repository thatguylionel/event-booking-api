package routes

import "github.com/gin-gonic/gin"

func EventsRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
