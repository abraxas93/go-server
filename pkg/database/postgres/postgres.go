package postgres

import (
	"database/sql"
	"fmt"
	"go-server/pkg/config"
	"go-server/pkg/utils/logger"

	_ "github.com/lib/pq"
)

// postgres://admin:abraxas93@localhost:5432/local?sslmode=disable
func Connect(c config.PostgresConfig) (*sql.DB, error) {
	log := logger.GetLogger()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
	log.Info(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close() // Close the connection before returning an error
		return nil, err
	}

	log.Info("Established a successful connection!")
	return db, nil
}
