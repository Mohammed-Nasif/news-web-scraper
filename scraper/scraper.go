package scraper

import (
	"log"
	"news-web-scraper/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeArticles() ([]models.Article, error) {
	var articles []models.Article
	urls := []string{"https://techcrunch.com/", "https://arstechnica.com/"} // Replace with real URLs

	for _, url := range urls {
		// Select the appropriate scraping function based on the URL
		var scrapeFunc func(*goquery.Document) []models.Article
		switch {
		case strings.Contains(url, "techcrunch"):
			scrapeFunc = scrapeTechCrunch
		case strings.Contains(url, "arstechnica"):
			scrapeFunc = scrapeArsTechnica
		default:
			log.Println("Unknown site:", url)
			continue
		}

		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Println("Error loading page:", url, err)
			continue
		}

		articles = append(articles, scrapeFunc(doc)...)
	}

	return articles, nil
}
