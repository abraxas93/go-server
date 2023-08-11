package user

import "time"

type IUserService interface {
}

type UserService struct {
	repo UserRepository
}

func NewService(r *UserRepository) {

}

func (s *UserService) CreateNewUser(name string, password string) {
	newUser := User{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUser.encryptPassword()
}
