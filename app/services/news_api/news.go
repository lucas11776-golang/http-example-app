package newsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/env"
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
