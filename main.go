package main

import (
	"log"
	"news-web-scraper/db"
	"news-web-scraper/routes"
)

func main() {
	db.ConnectDB()

	server := routes.SetupRouter()

	err := routes.ScrapeArticlesAndInsetToDB()

	if err != nil {
		log.Fatal("Error scraping articles:", err)
		return
	}

	log.Println("Server running on http://localhost:8080")
	server.Run(":8080")
}
