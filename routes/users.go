package routes

import (
	"net/http"
	"tgl/eventapi/models"
	"tgl/eventapi/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials", "error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	context.JSON(http.StatusOK, gin.H{"message": "User authenticated", "token": token})
}
