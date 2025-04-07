package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/middlewares"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Apply authentication middleware to all routes in this group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.GET("/events", getEvents)
	authenticated.GET("/events/:id", getEvent)

	server.POST("/signup", signupUser)
	server.POST("/login", loginUser)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	// Get the user ID from the context (set by the authentication middleware)
	userID := context.GetInt64("userID")
	var event models.Event
	// Bind JSON to the event struct
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userID == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Set the user ID from the token
	event.UserID = int(userID) // Set the user ID from the token

	err := models.Save(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userID := context.GetInt64("userID") // del middleware

	// Traes el evento existente
	existingEvent, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Validas que el usuario autenticado sea el creador
	if existingEvent.UserID != int(userID) {
		context.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this event"})
		return
	}

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = int(id)
	event.UserID = existingEvent.UserID

	err = models.UpdateEventByID(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	err = models.DeleteEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
