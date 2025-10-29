package routes

import (
	"net/http"
	"strconv"

	"example.com/evently-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

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

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered for event"})
}

func cancelRegisteration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

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

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	err = event.Unregister(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not unregister user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user unregistered for event"})
}
