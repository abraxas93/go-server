package user

import (
	"context"
	"go-server/pkg/database/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u User) (int64, error)
	FindByID(ctx context.Context, id int) (*models.User, error)
	DeleteByID(ctx context.Context, id int) error
}
