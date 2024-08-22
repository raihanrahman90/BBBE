package middleware

import (
	"bbbe/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		tokenSplit := strings.Split(tokenString, " ")
		token := tokenSplit[1] 

		claims, err := utils.ParseJWT(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.set("userId", claims.UserId)
		c.Set("username", claims.Username)
    	c.Set("access", claims.Access)
		c.Next()
	}
}
