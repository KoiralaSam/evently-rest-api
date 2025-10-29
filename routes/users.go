package routes

import (
	"net/http"

	"example.com/evently-rest-api/models"
	"example.com/evently-rest-api/utils"
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

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login sucessfull", "token": token})
}

// func getRegistration(ctx *gin.Context) {

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
// 		return
// 	}
// 	event, err := models.GetEventByID(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, event)
// }
