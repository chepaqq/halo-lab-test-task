package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s host=%s password=%s port=%s sslmode=%s", cfg.Username, cfg.DBName, cfg.Host, cfg.Password, cfg.Port, cfg.SSLMode),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
