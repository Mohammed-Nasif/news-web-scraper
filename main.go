package main

import (
	"log"
	"os"
	"time"
	"web-scraper/db"
	"web-scraper/routes"
	"web-scraper/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.ConnectDB()

	log.Println("Starting initial scraping...")
	services.ScrapeArticlesAndInsertToDB()

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("Starting rescraping...")
			err := services.ScrapeArticlesAndInsertToDB()
			if err != nil {
				log.Println("Error during rescraping:", err)
			} else {
				log.Println("Rescraping completed successfully.")
			}
		}
	}()

	server := routes.SetupRouter()
	port := os.Getenv("PORT")
	baseURL := os.Getenv("BASE_URL")
	log.Printf("Server running on %s:%s", baseURL, port)
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
