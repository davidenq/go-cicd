package infra

import (
	"github.com/davidenq/go-cicd/cmd/gocicd/infra/config"
	"github.com/davidenq/go-cicd/cmd/gocicd/interface/controllers"
	"github.com/davidenq/go-cicd/cmd/gocicd/middleware"
	"github.com/gin-gonic/gin"
)

//NewServer .
func NewServer() error {

	r := gin.Default()

	unAuthorizedRoutes := r.Group("/")
	unAuthorizedRoutes.GET("/health", controllers.Health)
	unAuthorizedRoutes.GET("/", controllers.Health)

	authorizedRoutes := r.Group("/DevOps")
	authorizedRoutes.Use(middleware.CheckAPIKEY())
	authorizedRoutes.Use(middleware.CheckUnAuthorizedMethods())
	authorizedRoutes.Any("/", controllers.Handler)

	err := r.Run(":" + config.PORT)
	if err != nil {
		return err
	}
	return nil
}
