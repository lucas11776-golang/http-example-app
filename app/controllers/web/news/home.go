package news

import (
	newsapi "server/app/services/news_api"
	"time"

	"github.com/lucas11776-golang/http"
)

// Comment
func Home(req *http.Request, res *http.Response) *http.Response {
	return res.View("index", http.ViewData{
		"articles": fetch(req.URL.Query().Get("q"), "", 50),
		"q":        req.URL.Query().Get("q"),
	})
}

// Comment
func Category(req *http.Request, res *http.Response) *http.Response {
	return res.View("category", http.ViewData{
		"category": req.Parameters.Get("category"),
		"articles": fetch(req.URL.Query().Get("q"), req.Parameters.Get("category"), 50),
		"q":        req.URL.Query().Get("q"),
	})
}

// Comment
func fetch(search string, category string, limit int) *[]newsapi.Article {
	articles, err := newsapi.TopHeadlines(search, category, limit, time.Now().Format("2006-01-02"))

	if err != nil {
		articles = make([]newsapi.Article, 0)
	}

	// articles := make([]newsapi.Article, 0)

	return &articles
}
