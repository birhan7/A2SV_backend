package controller

import (
	"net/http"
	"task-management/domain"
	"task-management/infrastructure"
	"task-management/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase *usecase.TaskUseCase
}

type UserController struct {
	userUseCase *usecase.UserUseCase
}

func NewTaskController(tc *usecase.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: tc,
	}
}

func NewUserController(uc *usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {
	var task domain.Task
	err := ctx.ShouldBind(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " Bad Request."})
		return
	}
	err = tc.taskUseCase.Create(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message:": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message:": " Task Created Successfully."})
}

func (tc *TaskController) FindTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tc.taskUseCase.FetchById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " Task with id" + id + " is not available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Task: ": task})
}

func (tc *TaskController) Tasks(ctx *gin.Context) {
	tasks, err := tc.taskUseCase.Fetch()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Tasks: ": tasks})
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	var task domain.Task
	id := ctx.Param("id")
	err := ctx.ShouldBind(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " Bad Request."})
		return
	}
	err = tc.taskUseCase.Update(id, task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message: ": "Task updated."})
}

func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := tc.taskUseCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message: ": "Task deleted."})
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " Bad Request."})
		return
	}
	err = uc.userUseCase.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error:": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message:": " User Created Successfully."})
}

func (uc *UserController) FinduserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.userUseCase.FetchById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " User with id" + id + " is not available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user: ": user})
}

func (uc *UserController) Users(ctx *gin.Context) {
	users, err := uc.userUseCase.Fetch()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users: ": users})
}

func (uc *UserController) Updateuser(ctx *gin.Context) {
	var user domain.User
	id := ctx.Param("id")
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message:": " Bad Request."})
		return
	}
	err = uc.userUseCase.Update(id, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message: ": "user updated."})
}

func (uc *UserController) Deleteuser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.userUseCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message: ": "user deleted."})
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	existingUser, _ := uc.userUseCase.FetchById(user.ID)
	if existingUser.Email == user.Email {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": "User already Registered."})
		return
	}
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	user.Password = hashedPassword
	err = uc.userUseCase.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message: ": "user registered."})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user domain.User
	id := ctx.Param("id")
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	existingUser, err := uc.userUseCase.FetchById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	verified := infrastructure.VerifiedPassword(existingUser.Password, user.Password) && existingUser.Email == user.Email
	if !verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": "Incorrect username or password."})
		return
	}
	token, err := infrastructure.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Token: ": token})
}
