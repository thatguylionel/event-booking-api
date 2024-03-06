package main

import (
	"tgl/eventapi/db"
	"tgl/eventapi/routes"

	"github.com/gin-gonic/gin"
)

const serverAddress = "localhost:8080"

func main() {
	db.InitDB()
	server := gin.Default()
	routes.EventsRouter(server)

	server.Run(serverAddress)
}
