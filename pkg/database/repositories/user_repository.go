package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"go-server/internal/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, u user.User) (int64, error) {
	query := `
		INSERT INTO users (name, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int64

	err := r.DB.QueryRowContext(ctx, query, u.Name, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	var user user.User
	query := "SELECT id, name, password, created_at, updated_at FROM users WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &user, err
	}
	fmt.Printf("%p\n", &user)
	fmt.Println("----")
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
