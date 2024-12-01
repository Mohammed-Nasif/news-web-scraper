package services

import (
	"log"
	"web-scraper/db"
	"web-scraper/scraper"
)

func ScrapeArticlesAndInsertToDB() error {
	articles, err := scraper.ScrapeArticles()
	if err != nil {
		log.Println("Error scraping articles: ", err)
		return err
	}

	for _, article := range articles {
		rowsAffected, err := db.InsertArticleToDB(article.Title, article.Link, article.Timestamp)
		if err != nil {
			log.Println("Error inserting article into database:", err)
			return err
		}

		if rowsAffected == 0 {
			log.Println("Skipped duplicate article:", article.Title)
		} else {
			log.Println("Article inserted:", article.Title)
		}
	}
	log.Println("Articles scraped and inserted successfully")

	return nil
}
