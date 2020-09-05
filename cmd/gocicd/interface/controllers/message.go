package controllers

import (
	"net/http"

	"github.com/davidenq/go-cicd/cmd/gocicd/domain"
	"github.com/davidenq/go-cicd/cmd/gocicd/types"
	"github.com/gin-gonic/gin"
)

//Handler .
func Handler(c *gin.Context) {
	var data types.BodyMessage
	c.BindJSON(&data)
	message := domain.GenerateMessage(&data)
	c.String(http.StatusOK, message)
}
