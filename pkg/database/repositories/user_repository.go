package repositories

import (
	"context"
	"database/sql"
	"go-server/pkg/database/models"
)

type UserRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, u models.User) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO users (name, password, created_at, updated_at) VALUES ($1, $2, $3, $4)", u.Name, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, password, created_at, updated_at FROM users WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
