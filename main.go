package main

import (
	"apex/controllers"
	"apex/core"
	"apex/repository"
	"apex/routes"
	"github.com/draco121/horizon/database"
	"github.com/draco121/horizon/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func RunApp() {
	utils.Logger.Info("Starting projectmanager service")
	client := database.NewMongoDatabase(os.Getenv("MONGODB_URI"))
	db := client.Database("projectmanager")
	repo := repository.NewProjectRepository(db)
	service := core.NewProjectService(client, repo)
	controller := controllers.NewControllers(service)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(utils.Logger.Out))
	routes.RegisterRoutes(controller, router)
	utils.Logger.Info("started projectmanager service")
	err := router.Run()
	if err != nil {
		utils.Logger.Fatal(err)
		return
	}
}
func main() {
	_ = godotenv.Load()
	RunApp()
}
