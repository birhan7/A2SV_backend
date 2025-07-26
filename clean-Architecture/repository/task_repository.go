package repository

import (
	"context"
	"errors"
	"task-management/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewTaskRepository(coll *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: coll,
		context:    context.Background(),
	}
}

func (tr *TaskRepository) Create(task domain.Task) error {
	_, err := tr.collection.InsertOne(tr.context, task)
	return err
}

func (tr *TaskRepository) Fetch() ([]domain.Task, error) {
	var tasks []domain.Task
	pointer, err := tr.collection.Find(tr.context, bson.D{{}})
	if err != nil {
		return tasks, err
	}
	if pointer.Next(tr.context) {
		var task domain.Task
		err = pointer.Decode(&task)
		if err == nil {
			tasks = append(tasks, task)
		}
	}
	if pointer.Err() != nil {
		return []domain.Task{}, err
	}
	return tasks, nil
}

func (tr *TaskRepository) FetchById(id string) (domain.Task, error) {
	var task domain.Task
	filter := bson.D{{Key: "_id", Value: id}}
	err := tr.collection.FindOne(tr.context, filter).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (tr *TaskRepository) Delete(id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := tr.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no task found with the given id")
	}
	return nil
}

func (tr *TaskRepository) Update(id string, task domain.Task) error {
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
		return nil
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: updateFields},
	}
	_, err := tr.collection.UpdateOne(tr.context, filter, update)

	if err != nil {
		return err
	}
	return nil
}
