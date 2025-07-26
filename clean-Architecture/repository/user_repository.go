package repository

import (
	"context"
	"errors"
	"task-management/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewUserRepository(coll *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: coll,
		context:    context.Background(),
	}
}

func (ur *UserRepository) Create(user domain.User) error {
	_, err := ur.collection.InsertOne(ur.context, user)
	return err
}

func (ur *UserRepository) Fetch() ([]domain.User, error) {
	var users []domain.User
	pointer, err := ur.collection.Find(ur.context, bson.D{{}})
	if err != nil {
		return users, err
	}
	if pointer.Next(ur.context) {
		var user domain.User
		err = pointer.Decode(user)
		if err == nil {
			users = append(users, user)
		}
	}
	if pointer.Err() != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (ur *UserRepository) FetchById(id string) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}
	err := ur.collection.FindOne(ur.context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Delete(id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := ur.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found with the given iD")
	}
	return nil
}

func (ur *UserRepository) Update(id string, user domain.User) error {
	updateFields := bson.D{}

	if user.Email != "" {
		updateFields = append(updateFields, bson.E{Key: "email", Value: user.Email})
	}
	if user.Password != "" {
		updateFields = append(updateFields, bson.E{Key: "password", Value: user.Password})
	}

	if len(updateFields) == 0 {
		return nil
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: updateFields},
	}
	_, err := ur.collection.UpdateOne(ur.context, filter, update)

	if err != nil {
		return err
	}
	return nil
}
