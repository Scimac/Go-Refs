package main

import (
	db "events-booking/db"
	"events-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // Start the server on port 8080 - domain auto decided by OS
}
