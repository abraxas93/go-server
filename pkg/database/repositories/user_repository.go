package repositories

import (
	"context"
	"database/sql"
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
		if err == sql.ErrNoRows {
			return nil, nil // User not found, return nil and nil error
		}
		return nil, err // Other error, return nil user and the error
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

func (r *UserRepository) FindAll(ctx context.Context) ([]user.User, error) {
	var users []user.User
	query := "SELECT * FROM users"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user.User
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		// No users found, return empty slice and nil error
		return []user.User{}, nil
	}

	return users, nil
}
