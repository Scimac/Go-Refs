package routes

import (
	events "events-booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllEvents(c *gin.Context) {
	e, err := events.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events. Please try again later.", "error": err.Error()})
		return
	}

	if len(e) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No events found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": e})
}

func getEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the id", "error": err.Error()})
	}
	e, err := events.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event. Please try again later.", "error": err.Error()})
		return
	}

	if e.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": e})
}

func createEvent(c *gin.Context) {
	var e events.Event

	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// This value was set from the middleware - Auth
	userId := c.GetInt64("userId")
	e.UserID = userId

	err = e.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event. Please try again later.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": e})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event ID", "error": err.Error()})
		return
	}

	event, err := events.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event.", "error": err.Error()})
		return
	}

	loggedInUserID := c.GetInt64("userId")
	if loggedInUserID != event.UserID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only event owners can update the info."})
		return
	}

	var updatedEvent events.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event ID", "error": err.Error()})
		return
	}

	e, err := events.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve the event.", "error": err.Error()})
		return
	}

	loggedInUserID := c.GetInt64("userId")
	if loggedInUserID != e.UserID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only event owners can delete the info."})
		return
	}

	err = e.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event deleted successfully", "event": e})
}
