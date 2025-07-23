package router

import (
	"task-manager-api/controllers"
	"task-manager-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	controller := controllers.CreateController()

	Router.GET("/tasks", controller.GetAllTasks)
	Router.GET("/tasks/:id", controller.GetTaskByID)
	Router.PUT("/tasks/:id", controller.UpdateTask)
	Router.DELETE("/tasks/:id", middleware.AuthMiddleware(), controller.DeleteTask)
	Router.POST("/tasks", middleware.AuthMiddleware(), controller.CreateTask)
	Router.POST("/register", controller.Register)
	Router.POST("/login", controller.Login)

	return Router
}
