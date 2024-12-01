package main

import (
	"log"
	"time"
	"web-scraper/db"
	"web-scraper/routes"
	"web-scraper/services"
)

func main() {
	db.ConnectDB()

	log.Println("Starting initial scraping...")
	services.ScrapeArticlesAndInsertToDB()

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("Starting Rescraping...")
			err := services.ScrapeArticlesAndInsertToDB()
			if err != nil {
				log.Println("Error during Rescraping:", err)
			} else {
				log.Println("Rescraping completed successfully.")
			}
		}
	}()

	server := routes.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: ", err)
	}
}
