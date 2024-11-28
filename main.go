package main

import (
	"log"
	"news-web-scraper/db"
	"news-web-scraper/routes"
	"time"
)

func main() {
	db.ConnectDB()

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("Starting Rescraping...")
			err := routes.ScrapeArticlesAndInsetToDB()
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
