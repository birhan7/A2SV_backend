package router

import (
	"context"
	"fmt"
	"log"
	"task-management/Delivery/controller"
	"task-management/infrastructure"
	"task-management/repository"
	"task-management/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetUp(databaseClient *mongo.Client) {
	taskCollection := databaseClient.Database("Tasks_storage").Collection("Tasks")
	userCollection := databaseClient.Database("Users_storage").Collection("Users")
	taskRepo := repository.NewTaskRepository(taskCollection)
	userRepo := repository.NewUserRepository(userCollection)
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	taskController := controller.NewTaskController(taskUseCase)
	userController := controller.NewUserController(userUseCase)

	router := gin.Default()
	protectedRoute := router.Group("")
	publicRoute := router.Group("")
	protectedRoute.Use(infrastructure.AuthMiddleware())

	protectedRoute.GET("/tasks", taskController.Tasks)
	protectedRoute.GET("/tasks/:id", taskController.FindTaskById)
	protectedRoute.POST("/tasks", taskController.CreateTask)
	protectedRoute.DELETE("/tasks/:id", taskController.DeleteTask)
	protectedRoute.PATCH("/tasks/:id", taskController.UpdateTask)
	publicRoute.POST("/register", userController.Register)
	publicRoute.POST("/login/:id", userController.Login)

	router.Run(":9000")
}

func InitDB(address string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(address)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongoDB successfully.")
	return client

}

func CloseDB(database *mongo.Client) {
	if err := database.Disconnect(context.TODO()); err != nil {
		log.Fatalln("Error closing the database connection.")
		return
	}
	fmt.Println("Database closed successfully.")
}
