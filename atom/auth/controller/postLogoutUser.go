package atom_auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PostLogoutUser(ctx *gin.Context) {

	_, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Already logged out"})
		return
	}

	isSecure := os.Getenv("IS_PROD") == "true"

	if !isSecure && ctx.Request.TLS != nil {
		isSecure = true
	}

	// utils_token.InvalidateRefreshToken(refreshTokenString)

	//remove cookie
	ctx.SetCookie("refresh_token", "", -1, "/", "", isSecure, true)

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
