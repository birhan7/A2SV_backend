package data

import (
	"task-manager-api/models"
	"time"
)

// Service structure that interacts with the Repository Layer. For the time being it uses in memory storage.
type Service struct {
	repository []models.Task
}

// Mock data that represents repository layer
var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func CreateService() Service {
	service := Service{}
	service.repository = tasks
	return service
}

func (s *Service) GetTasks() []models.Task {
	return tasks
}

func (s *Service) GetTask(id string) *models.Task {
	for key, val := range tasks {
		if val.ID == id {
			return &tasks[key]
		}
	}
	return nil
}

func (s *Service) RemoveTask(id string) bool {
	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Service) AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func (s *Service) UpdateTask(id string, task models.Task) bool {
	for key, value := range tasks {
		if id == value.ID {
			if task.Title != "" {
				tasks[key].Title = task.Title
			}
			if task.Description != "" {
				tasks[key].Description = task.Description
			}
			if task.Status != "" {
				tasks[key].Status = task.Status
			}
			return true
		}
	}
	return false
}
