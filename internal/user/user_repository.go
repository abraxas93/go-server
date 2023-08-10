package user

import (
	"context"
	"go-server/pkg/database/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, password string) error
	FindByID(ctx context.Context, id int) (*models.User, error)
	DeleteByID(ctx context.Context, id int) error
}
