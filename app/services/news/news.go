package news

import "server/models"

// Comment
func NewsExists(url string) bool {
	return true
}

// Comment
func NewsByUrl(url string) []models.Article {
	return []models.Article{}
}

// Comment
func NewsSave(url string, articles []models.Article) bool {
	return true
}
