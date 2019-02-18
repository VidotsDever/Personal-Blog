package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("authorization")
		if tokenStr == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not Authenticated",
			})
			c.Abort()
			return
		}
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusForbidden, gin.H{
					"message": "Not Authenticated",
				})
				return nil, fmt.Errorf("Not Authenticated")
			}
			return []byte("i-love-you"), nil
		})
		if !token.Valid {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not Authenticated",
			})
			c.Abort()
			return
		}
	}
}
