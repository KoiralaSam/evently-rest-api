package routes

import (
	"net/http"
	"strconv"

	"example.com/evently-rest-api/models"
	"example.com/evently-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	_, err = models.GetEventByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event sucessfully updated"})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"message": "could not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event sucessfully deleted"})
}
