package infra

import (
	"github.com/davidenq/go-cicd/cmd/gocicd/infra/config"
	"github.com/davidenq/go-cicd/cmd/gocicd/interface/controllers"
	"github.com/gin-gonic/gin"
)

//NewServer .
func NewServer() error {

	r := gin.Default()
	r.GET("/health", controllers.Health)
	r.POST("/DevOps", controllers.Handler)

	err := r.Run(":" + config.PORT)
	if err != nil {
		return err
	}
	return nil
}
