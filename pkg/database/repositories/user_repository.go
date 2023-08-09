package repositories

import "database/sql"

type UserRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
