package scraper

import (
	"log"
	"news-web-scraper/models"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeArticles() ([]models.Article, error) {
	var articles []models.Article
	urls := []string{"https://techcrunch.com/", "https://arstechnica.com/", "https://www.theverge.com/"}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			var scrapeFunc func(*goquery.Document) []models.Article
			switch {
			case strings.Contains(url, "techcrunch"):
				scrapeFunc = scrapeTechCrunch
			case strings.Contains(url, "arstechnica"):
				scrapeFunc = scrapeArsTechnica
			case strings.Contains(url, "theverge"):
				scrapeFunc = scrapeTheverge
			default:
				log.Println("Unknown site:", url)
				return
			}

			doc, err := goquery.NewDocument(url)
			if err != nil {
				log.Println("Error loading page:", url, err)
				return
			}

			scrapedArticles := scrapeFunc(doc)
			mu.Lock()
			articles = append(articles, scrapedArticles...)
			mu.Unlock()
		}(url)
	}

	wg.Wait()
	return articles, nil
}
