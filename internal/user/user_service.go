package user

import (
	"context"
	"time"
)

type UserServiceInterface interface {
	CreateNewUser(name string, password string) (int64, error)
	GetUser(id int) (User, error)
}

type UserService struct {
	Repo UserRepositoryIface
}

func NewService(r UserRepositoryIface) UserService {
	return UserService{Repo: r}
}

func (s *UserService) CreateNewUser(name string, password string) (int64, error) {
	newUser := User{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUser.EncryptPassword()
	id, err := s.Repo.CreateUser(context.Background(), newUser)
	return id, err
}

func (s *UserService) GetUser(id int) (*User, error) {
	user, err := s.Repo.FindByID(context.Background(), id)
	return user, err
}
