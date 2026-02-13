package atom_auth

import (
	"errors"
	"log"
	"net/http"
	atom_auth "optimach_service/atom/auth"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	log.Println("[atom][auth][refreshToken] Receiving Refresh Token Controller request")

	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("[atom][auth][refreshToken] Invalid request body:", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
			"error":   "refresh_token is required in the body",
		})
		return
	}

	refreshTokenString := requestBody.RefreshToken

	if refreshTokenString == "" {
		log.Println("[atom][auth][refreshToken] Refresh token is empty.")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Unauthorized",
			"error":   "Refresh token not found, need login",
		})
		return
	}

	log.Println("refresh token string found in body")

	newAccessTokenString, newAccessExp, generateAccErr := atom_auth.RefreshAccessTokenUseCase(refreshTokenString)

	if generateAccErr != nil {
		log.Printf("[atom][auth][controller][RefreshToken] Refresh error: %s", generateAccErr.Error())
		if errors.Is(generateAccErr, atom_auth.ErrValidationFailed) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "Unauthorized",
				"error":   generateAccErr.Error(),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "Internal Server Error",
				"error":   "Server error during token refresh",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessTokenString,
		"expires_at":   newAccessExp,
		"status":       200,
		"message":      "Access token refreshed successfully",
	})
}
