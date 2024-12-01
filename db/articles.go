package db

import (
	"time"
	"web-scraper/models"
)

func SelectArticlesFromDB() ([]models.Article, error) {
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

func InsertArticleToDB(title string, link string, timestamp time.Time) (int64, error) {
	insertQuery := `
    INSERT INTO articles (title, link, timestamp) 
    VALUES ($1, $2, $3)
    ON CONFLICT (link) DO NOTHING
    `

	result, err := DB.Exec(insertQuery, title, link, timestamp)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
