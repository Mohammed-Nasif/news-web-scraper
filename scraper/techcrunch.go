package scraper

import (
	"log"
	"time"
	"web-scraper/models"

	"github.com/PuerkitoBio/goquery"
)

func scrapeTechCrunch(doc *goquery.Document) []models.Article {
	var articles []models.Article

	doc.Find(".loop-card").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		timestampStr, _ := s.Find("time").Attr("datetime")
		timestamp, _ := time.Parse(time.RFC3339, timestampStr)

		if title == "" || link == "" {
			log.Println("Skipping article due to missing title or link")
			return
		}

		articles = append(articles, models.Article{
			Title:     title,
			Link:      link,
			Timestamp: timestamp,
		})
	})

	return articles
}
