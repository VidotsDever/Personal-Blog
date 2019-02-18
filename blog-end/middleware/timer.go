package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RecordTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		fmt.Println("Request path ", c.Request.URL.Path, "Duration ", time.Since(startTime))
	}
}
