package routes

import (
	events "events-booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerToEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event ID", "error": err.Error()})
		return
	}

	event, err := events.GetEventByID(eventid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event.", "error": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for the event."})
}

func deleteRegisteration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event ID", "error": err.Error()})
		return
	}

	event, err := events.GetEventByID(eventid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event.", "error": err.Error()})
		return
	}

	err = event.DeleteRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the registration for the event.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the registration for the event."})

}

func getAllRegistrations(c *gin.Context) {
	regs, err := events.GetAllRegistrations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"registrations": regs})
}
