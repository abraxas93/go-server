package user

import "go-server/pkg/database/models"

type UserRepository interface {
	CreateUser(name string, password string) error
	FindByID(id int) (*models.User, error)
}
