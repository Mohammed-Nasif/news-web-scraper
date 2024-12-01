package services

import (
	"log"
	"net/http"
	"web-scraper/db"

	"github.com/gin-gonic/gin"
)

func GetArticles(context *gin.Context) {
	articles, err := db.SelectArticlesFromDB()
	if err != nil {
		log.Println("Error Getting articles: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not Get articles"})
		return
	}
	context.JSON(http.StatusOK, articles)
}
