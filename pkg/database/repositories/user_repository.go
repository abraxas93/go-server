package repositories

import "database/sql"

type UserRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(name string, password string) error {
	_, err := r.DB.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", name, password)
	if err != nil {
		return err
	}
	return nil
}
