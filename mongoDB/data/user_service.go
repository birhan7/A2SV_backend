package data

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"task-manager-api/middleware"
	"task-manager-api/models"
)

var users = make(map[string]*models.User)

func (s *Service) RegisterUser(user models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	users[user.Email] = &user

	return nil
}

func (s *Service) UserLogin(user models.User) (string, error) {
	existingUser, ok := users[user.Email]
	if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("error: Invalid email or password")
	}
	token, err := middleware.CreateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
