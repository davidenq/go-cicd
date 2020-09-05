package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CheckUnAuthorizedMethods .
func CheckUnAuthorizedMethods() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		switch method {
		case "POST":
			c.Next()
		default:
			c.String(http.StatusMethodNotAllowed, "ERROR")
			c.Abort()
		}
	}
}
