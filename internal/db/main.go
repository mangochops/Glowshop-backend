package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	// Option 1: Use DB_HOST as a full connection string if provided
	connStr := os.Getenv("DB_HOST")
	if connStr != "" {
		return sql.Open("pgx", connStr)
	}

	// Option 2: Build connection string from individual vars
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	return sql.Open("pgx", connStr)
}
