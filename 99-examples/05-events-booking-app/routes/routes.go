package routes

import (
	"events-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)     // Endpoint to get all events
	server.GET("/events/:id", getEventByID) // Endpoint to get a specific event by ID

	// Gin allows multiple handlers and are executed from left to right
	// But not a great way if same middleware used for multiple routes
	// server.POST("/events", middlewares.Authenticate, createEvent) // Endpoint to create a new event

	// relative path can be given
	authBasedApis := server.Group("/")
	authBasedApis.Use(middlewares.Authenticate)

	// now the authBasedApis group is ued to listen to these paths
	authBasedApis.POST("/events", createEvent)       // Endpoint to create a new event
	authBasedApis.PUT("/events/:id", updateEvent)    // Endpoint to update an event
	authBasedApis.DELETE("/events/:id", deleteEvent) // Endpoint to delete an event

	authBasedApis.POST("/events/:id/register", registerToEvent)       // Endpoint to register for an event
	authBasedApis.DELETE("/events/:id/register", deleteRegisteration) // endpoint to cancel the registration
	server.GET("/registrations", getAllRegistrations)

	server.GET("/users", getAllUsers)
	server.POST("/signup", userSignup) // Endpoint to sign up for users
	server.POST("/login", userLogin)   // Endpoint  to login the user
}
