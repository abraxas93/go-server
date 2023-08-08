package postgres

import (
	"database/sql"
	"fmt"
	"go-server/pkg/config"

	_ "github.com/lib/pq"
)

func Connect(c config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close() // Close the connection before returning an error
		return nil, err
	}

	fmt.Println("Established a successful connection!")
	return db, nil
}
