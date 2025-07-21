package data

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"task-manager-api/models"
	"time"
)

// Service structure that interacts with the Repository Layer. For the time being it uses in memory storage.
type Service struct {
	repository *mongo.Database
	collection mongo.Collection
}

// Mock data that represents repository layer
var tasks = []interface{}{
	models.Task{ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5)},
	models.Task{ID: "2", Status: "Pending", Title: "Design API Endpoints", Description: "Plan REST API structure and routes for task management", DueDate: time.Now().AddDate(0, 0, 2)},
	models.Task{ID: "3", Status: "Completed", Title: "Setup Database", Description: "Initialize PostgreSQL and create necessary tables", DueDate: time.Now().AddDate(0, 0, -3)},
	models.Task{ID: "4", Status: "In progress", Title: "Write Unit Tests", Description: "Cover task controller with unit tests using testify", DueDate: time.Now().AddDate(0, 0, 7)},
	models.Task{ID: "5", Status: "Pending", Title: "Dockerize App", Description: "Create Dockerfile and docker-compose for development", DueDate: time.Now().AddDate(0, 0, 10)},
}

func ConnectDB() *mongo.Database {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB successfully.")
	collection := client.Database("Tasks_storage").Collection("Tasks")
	collection.DeleteMany(context.TODO(), bson.D{{}})
	collection.InsertMany(context.TODO(), tasks)
	return collection.Database()

}
func CreateService() Service {
	var service Service
	service.repository = ConnectDB()
	service.collection = *service.repository.Collection("Tasks")
	return service
}

func (s *Service) CloseDB() {
	err := s.repository.Client().Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from MongoDB")
}

func (s *Service) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	pointer, err := s.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return tasks, err
	}
	for pointer.Next(context.TODO()) {
		var task models.Task
		err = pointer.Decode(&task)
		tasks = append(tasks, task)
	}
	if pointer.Err() != nil {
		return []models.Task{}, err
	}
	return tasks, nil
}

func (s *Service) GetTask(id string) *models.Task {
	filter := bson.D{{Key: "id", Value: id}}
	var task models.Task
	err := s.collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		log.Println("Error finding task by ID:", err)
		return nil
	}
	return &task
}

func (s *Service) RemoveTask(id string) error {
	filter := bson.D{{Key: "id", Value: id}}
	result, err := s.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting task:", err)
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("No task found with the given ID")
	}

	log.Println("Task deleted successfully")
	return nil
}

func (s *Service) AddTask(task models.Task) error {
	_, err := s.repository.Collection("Tasks").InsertOne(context.TODO(), task)
	return err
}

func (s *Service) UpdateTask(id string, task models.Task) error {
	updateFields := bson.D{}

	if task.Title != "" {
		updateFields = append(updateFields, bson.E{Key: "title", Value: task.Title})
	}
	if task.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: task.Description})
	}
	if !task.DueDate.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "due_date", Value: task.DueDate})
	}
	if task.Status != "" {
		updateFields = append(updateFields, bson.E{Key: "status", Value: task.Status})
	}

	if len(updateFields) == 0 {
		log.Println("No fields to update")
		return nil
	}

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: updateFields},
	}
	_, err := s.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println("Error updating task:", err)
		return err
	}
	return nil
}
