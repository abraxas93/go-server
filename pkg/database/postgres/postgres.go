package postgres

import (
	"database/sql"
	"fmt"
	"go-server/pkg/config"
	"go-server/pkg/logger"

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
	log := logger.GetLogger()
	log.Info("Established a successful connection!")
	return db, nil
}
