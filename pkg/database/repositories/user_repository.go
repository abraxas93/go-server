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

func (r *UserRepository) CreateUser(ctx context.Context, name string, password string) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO users (name, password) VALUES ($1, $2)", name, password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, password, created_at, updated_at FROM users WHERE id = $1"
	row := ur.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
