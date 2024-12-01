package routes

import (
	"log"
	"net/http"
	"news-web-scraper/db"
	"news-web-scraper/middlewares"
	"news-web-scraper/scraper"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.RateLimiter())
	router.GET("/articles", fetchArticles)

	return router
}

func fetchArticles(context *gin.Context) {
	articles, err := db.GetAllArticles()
	if err != nil {
		log.Println("Error fetching articles: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch articles"})
		return
	}
	context.JSON(http.StatusOK, articles)
}

func ScrapeArticlesAndInsetToDB() error {
	articles, err := scraper.ScrapeArticles()
	if err != nil {
		log.Println("Error scraping articles: ", err)
		return err
	}

	for _, article := range articles {
		err := db.InsertArticle(article.Title, article.Link, article.Timestamp)
		if err != nil {
			log.Println("Error inserting article:", article.Title, err)
			return err
		}
	}
	log.Println(http.StatusOK, "Scraping and insertion completed successfully")

	return nil
}
