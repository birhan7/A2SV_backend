package router

import (
	"github.com/gin-gonic/gin"
	"task-manager-api/controllers"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	controller := controllers.CreateController()

	Router.GET("/tasks", controller.GetAllTasks)
	Router.GET("/tasks/:id", controller.GetTaskByID)
	Router.PUT("/tasks/:id", controller.UpdateTask)
	Router.DELETE("/tasks/:id", controller.DeleteTask)
	Router.POST("/tasks", controller.CreateTask)

	return Router
}
