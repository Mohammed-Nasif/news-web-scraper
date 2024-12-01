package controllers

import (
	"web-scraper/services"

	"github.com/gin-gonic/gin"
)

func GetArticles(context *gin.Context) {
	services.GetArticles(context)
}
