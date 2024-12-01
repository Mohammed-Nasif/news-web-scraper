package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = createArticlesTable()
	if err != nil {
		log.Fatalf("Error creating articles table: %v", err)
	}

	log.Println("Connected to the database.")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func createArticlesTable() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS articles (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255),
			link TEXT UNIQUE,
			timestamp TIMESTAMP
		)`

	_, err := DB.Exec(createTableQuery)
	return err
}
