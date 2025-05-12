package newsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/env"
	"server/models"
	"strings"
	"time"
)

type ArticleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Article     string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt time.Time     `json:"publishedAt"`
	Content     string        `json:"content"`
}

type News struct {
	Status       string    `json:"status"`
	TotalResults int64     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Comment
func Map(articles []Article, callback func(article Article) models.Article) []models.Article {
	tranformed := []models.Article{}

	for _, article := range articles {
		tranformed = append(tranformed, callback(article))
	}

	return tranformed
}

// Comment
func Fetch(search string, category string, limit int) *[]models.Article {
	articles, err := TopHeadlines(search, category, limit, time.Now().Format("2006-01-02"))

	if err != nil {
		articles = make([]Article, 0)
	}

	transformed := Map(articles, func(article Article) models.Article {
		return models.Article{
			Publisher:   article.Source.Name,
			PublishedAt: article.PublishedAt,
			Image:       article.UrlToImage,
			Title:       article.Title,
			Description: article.Description,
			Content:     article.Content,
			Url:         article.Url,
		}
	})

	return &transformed
}

// Comment
func TopHeadlines(search string, category string, limit int, from string) ([]Article, error) {
	news, err := Request(Url("top-headlines", search, category, limit, from))

	if err != nil {
		return make([]Article, 0), err
	}

	return news.Articles, nil
}

// Comment
func Url(topic string, search string, category string, limit int, from string) string {
	query := []string{}

	if search != "" {
		query = append(query, fmt.Sprintf("q=%s", search))
	}

	if category != "" {
		query = append(query, fmt.Sprintf("category=%s", category))
	}

	query = append(query, fmt.Sprintf("pageSize=%d", limit))
	query = append(query, fmt.Sprintf("from=%s", from))
	query = append(query, fmt.Sprintf("apiKey=%s", env.Env("NEWS_API_KEY")))
	query = append(query, fmt.Sprintf("language=%s", "en"))

	return fmt.Sprintf("%s/%s?%s", env.Env("NEWS_API_URL"), topic, strings.Join(query, "&"))
}

// Comment
func Request(url string) (*News, error) {
	request, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		body = []byte{}
	}

	var news News

	err = json.Unmarshal(body, &news)

	if err != nil {
		return nil, err
	}

	return &news, nil
}
