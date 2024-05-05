package routes

import (
	"apex/controllers"
	"github.com/draco121/horizon/constants"
	"github.com/draco121/horizon/middlewares"
	"github.com/draco121/horizon/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(controllers controllers.Controllers, router *gin.Engine) {
	utils.Logger.Info("Registering routes...")
	v1 := router.Group("/v1")
	v1.POST("/project", middlewares.AuthMiddleware(constants.Write), controllers.CreateProject)
	v1.GET("/project", middlewares.AuthMiddleware(constants.Read), controllers.GetProject)
	v1.DELETE("/project", middlewares.AuthMiddleware(constants.Write), controllers.DeleteProject)
	utils.Logger.Info("Registered routes...")
}
