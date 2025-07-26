package usecase

import (
	"errors"
	"task-management/domain"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uc *UserUseCase) Create(user domain.User) error {
	if user.Password == "" || user.Email == "" {
		return errors.New("error: A user should have an email and password")
	}
	err := uc.userRepository.Create(user)
	return err
}

func (uc *UserUseCase) Fetch() ([]domain.User, error) {
	users, err := uc.userRepository.Fetch()
	return users, err
}

func (uc *UserUseCase) FetchById(id string) (domain.User, error) {
	user, err := uc.userRepository.FetchById(id)
	return user, err
}

func (uc *UserUseCase) Delete(id string) error {
	err := uc.userRepository.Delete(id)
	return err
}

func (uc *UserUseCase) Update(id string, user domain.User) error {
	err := uc.userRepository.Update(id, user)
	return err
}
