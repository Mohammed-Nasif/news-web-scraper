package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=nasif password=admin1234 dbname=newsdb sslmode=disable")

	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Connected to the database.")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createArticlesTable()
}

func createArticlesTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS articles (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		link TEXT,
		timestamp TIMESTAMP
	)`

	_, err := DB.Exec(createTableQuery)

	if err != nil {
		log.Fatal(`Error creating "articles" table: `, err)
	}

	log.Println("Articles table ensured.")
}
