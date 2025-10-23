package routes

import (
	"net/http"

	"example.com/evently-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Sucessfully created user"})

}
