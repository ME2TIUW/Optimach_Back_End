package middleware

import (
	"net/http"
	utils_token "optimach_service/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthenticateMiddleware() gin.HandlerFunc{
	return func (c *gin.Context) {
		tokenString := utils_token.ExtractToken(c)
		
		if tokenString == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status" : 401,
				"message" : "Unauthorized",
				"error" : "Access token required",
			})
			c.Abort()
			return
		}

		claims, err := utils_token.ValidateAccessToken(tokenString)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status" : 401,
				"message" : "Unauthorized",
				"error" : err.Error(),
			})

			c.Abort()
			return 
		}

		c.Set("userId", claims.Id_User)
		c.Set("accessClaims", claims)
		c.Next()

	}
}