package user

import (
	"context"
)

type UserRepositoryIface interface {
	CreateUser(ctx context.Context, u User) (int64, error)
	FindByID(ctx context.Context, id int) (*User, error)
	DeleteByID(ctx context.Context, id int) error
}
