package usecase

import (
	"errors"
	"task-management/domain"
)

type TaskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepository: taskRepo,
	}
}

func (uc *TaskUseCase) Create(task domain.Task) error {
	if task.Title == "" || task.Description == "" {
		return errors.New("error: A task should have a title and description")
	}
	err := uc.taskRepository.Create(task)
	return err
}

func (uc *TaskUseCase) Fetch() ([]domain.Task, error) {
	tasks, err := uc.taskRepository.Fetch()
	return tasks, err
}

func (uc *TaskUseCase) FetchById(id string) (domain.Task, error) {
	task, err := uc.taskRepository.FetchById(id)
	return task, err
}

func (uc *TaskUseCase) Delete(id string) error {
	err := uc.taskRepository.Delete(id)
	return err
}

func (uc *TaskUseCase) Update(id string, task domain.Task) error {
	err := uc.taskRepository.Update(id, task)
	return err
}
