package routes

import (
	"github.com/draco121/common/constants"
	"github.com/draco121/common/middlewares"
	"github.com/draco121/common/utils"
	"github.com/draco121/projectmanagerservice/controllers"
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
