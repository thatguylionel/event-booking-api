package routes

import (
	"tgl/eventapi/middlewares"

	"github.com/gin-gonic/gin"
)

func EventsRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	// Authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	//	Publically accessible routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
