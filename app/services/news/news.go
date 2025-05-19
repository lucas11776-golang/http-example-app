package news

import (
	"server/models"
	"server/utils/slices"

	"github.com/lucas11776-golang/orm"
)

// Comment
func NewsExists(url string) bool {
	count, err := orm.Model(models.NewsQuery{}).
		Where("url", "=", url).
		Count() // TODO: Added datetime to check if created at is less then 30 minutes.

	if err != nil {
		return false
	}

	return count > 0
}

// Comment
func transformNewsQueryArticles(articles []*models.NewsQueryArticles) []models.Article {
	return slices.Map(articles, func(article *models.NewsQueryArticles) models.Article {
		return models.Article{
			ID:          article.ID,
			Publisher:   article.Publisher,
			PublishedAt: article.PublishedAt,
			Image:       article.Image,
			Title:       article.Title,
			Description: article.Description,
			Content:     article.Content,
			Url:         article.Url,
		}
	})
}

// Comment
func NewsByUrl(url string) []models.Article {
	articles, err := orm.Model(models.NewsQueryArticles{}).
		Where("news_queries.url", "=", url).
		Join("articles", "news_queries.id", "=", "articles.news_query_id").
		OrderBy("articles.id", orm.ASC).
		Get()

	if err != nil {
		return []models.Article{}
	}

	return transformNewsQueryArticles(articles)
}

// Comment
func NewsSave(url string, articles []models.Article) error {
	news, err := orm.Model(models.NewsQuery{}).
		Insert(orm.Values{"url": url})

	if err != nil {
		return err
	}

	for _, article := range articles {
		orm.Model(models.Article{}).
			Insert(orm.Values{
				"news_query_id": news.ID,
				"publisher":     article.Publisher,
				"published_at":  article.PublishedAt,
				"image":         article.Image,
				"title":         article.Title,
				"description":   article.Description,
				"content":       article.Content,
				"url":           article.Url,
			})
	}

	return nil
}
