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

func (c *TaskController) DestroyController() {
	c.service.CloseDB()
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.service.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task := c.service.GetTask(id)
	if task != nil {
		ctx.JSON(http.StatusOK, task)
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

	err := c.service.AddTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
		return
	}

	err := c.service.UpdateTask(id, NewTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.RemoveTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

func (c *TaskController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.RegisterUser(user); err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, gin.H{"message": "User registered successfully"})
}

func (c *TaskController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.service.UserLogin(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Token": token})

}
