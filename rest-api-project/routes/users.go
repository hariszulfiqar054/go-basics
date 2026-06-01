package routes

import (
	"net/http"

	"example.com/rest_api/models"
	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid user data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid user data"})
		return
	}

	existingUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Could not find user"})
		return
	}
	if existingUser == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "message": "User not found"})
		return
	}

	err = utils.CheckPassword(user.Password, existingUser.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "message": "Password is incorrect"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
