package routes

import (
	"web-scraper/controllers"
	"web-scraper/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.RateLimiter())
	router.GET("/articles", controllers.GetArticles)

	return router
}
