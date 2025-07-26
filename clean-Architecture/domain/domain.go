package domain

import (
	"time"
)

type User struct {
	ID       string `bson:"_id" json:"id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Task struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
}

type TaskUseCase interface {
	Create(Task) error
	Delete(string) error
	Update(string, Task) error
	Fetch() ([]Task, error)
	FetchById(string) (Task, error)
}

type TaskRepository interface {
	Create(Task) error
	Delete(string) error
	Update(string, Task) error
	Fetch() ([]Task, error)
	FetchById(string) (Task, error)
}

type UserRepository interface {
	Create(User) error
	Delete(string) error
	Update(string, User) error
	Fetch() ([]User, error)
	FetchById(string) (User, error)
}

type UserUseCase interface {
	Create(User) error
	Delete(string) error
	Update(string) error
	Fetch() ([]User, error)
	FetchById(string) (User, error)
}
