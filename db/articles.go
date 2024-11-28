package db

import (
	"log"
	"news-web-scraper/models"
	"time"
)

func GetAllArticles() ([]models.Article, error) {
	rows, err := DB.Query("SELECT id, title, link, timestamp FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Link, &article.Timestamp)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func InsertArticle(title string, link string, timestamp time.Time) error {
	insertQuery := `
	INSERT INTO articles (title, link, timestamp) 
	VALUES ($1, $2, $3)
	ON CONFLICT (link) DO NOTHING
	`

	result, err := DB.Exec(insertQuery, title, link, timestamp)
	if err != nil {
		log.Println("Error inserting article:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("Skipped duplicate article:", title)
	} else {
		log.Println("Article inserted:", title)
	}

	return nil
}
