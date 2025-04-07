package middlewares

import (
	"net/http"

	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	claims, err := util.ValidateJWT(token)      // Print claims and error for debugging
	userId := int64(claims["userId"].(float64)) // Extract user ID from claims
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	context.Set("userID", userId) // Store user ID in context for later use
	context.Next()
}
