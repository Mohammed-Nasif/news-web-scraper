package scraper

import (
	"news-web-scraper/models"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func scrapeTechCrunch(doc *goquery.Document) []models.Article {
	var articles []models.Article

	doc.Find(".loop-card").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		timestamp := time.Now()

		articles = append(articles, models.Article{
			Title:     title,
			Link:      link,
			Timestamp: timestamp,
		})
	})

	return articles
}
