package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const key = "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c"

//CheckAPIKEY .
func CheckAPIKEY() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Parse-REST-API-Key")
		fmt.Println(apiKey)
		if apiKey == "" || apiKey != key {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
