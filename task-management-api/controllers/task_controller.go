package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-manager-api/data"
	"task-manager-api/models"
)

// Task controller interacts with the service layer that holds business logic
type TaskController struct {
	service data.Service
}

func CreateController() TaskController {
	controller := TaskController{}
	controller.service = data.CreateService()
	return controller
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks := c.service.GetTasks()
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task := c.service.GetTask(id)
	if task != nil {
		ctx.JSON(http.StatusOK, *task)
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newTask.Title == "" || newTask.Status == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task Must have title and status."})
		return
	}
	c.service.AddTask(newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var NewTask models.Task

	if err := ctx.BindJSON(&NewTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if NewTask.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Task Title should not be Empty."})
	}
	updated := c.service.UpdateTask(id, NewTask)
	if !updated {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	deleted := c.service.RemoveTask(id)
	if !deleted {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
