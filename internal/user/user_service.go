package user

import (
	"context"
	"time"
)

type IUserService interface {
}

type UserService struct {
	Repo UserRepository
}

func NewService(r UserRepository) *UserService {
	return &UserService{Repo: r}
}

func (s *UserService) CreateNewUser(name string, password string) error {
	newUser := User{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUser.encryptPassword()
	err := s.Repo.CreateUser(context.Background(), newUser)
	return err
}
