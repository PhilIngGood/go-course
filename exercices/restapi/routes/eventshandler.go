package routes

import (
	"net/http"
	"strconv"

	"course.go/restapi/models"
	"course.go/restapi/utils"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ID to int"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch event by id"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context) {
	// HTTP response
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvents(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to perform this action", "error": "Empty token"})
		return
	}

	userid, err := utils.ValidateJWT(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to perform this action", "error": "Invalid token"})
		return
	}
	var event models.Event
	err = c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err.Error()})
		return
	}

	event.UserID = userid
	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ID to int"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to fetch event by id"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ID to int"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to find event"})
		return
	}

	if event.DeleteEvent() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted ", "event": event})
}
