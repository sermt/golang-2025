package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func signupUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.SaveUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func loginUser(context *gin.Context) {
	var loginInput models.User
	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.ValidateCredentials(loginInput)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := util.GenerateJWT(user.Email, int64(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
