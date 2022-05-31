package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	user := os.Getenv("user")
	password := os.Getenv("password")
	database := os.Getenv("database")

	dsn := user + ":" + password + "@/" + database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
