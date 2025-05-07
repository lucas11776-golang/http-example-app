package news

import (
	newsapi "server/app/Services/news_api"
	"time"

	"github.com/lucas11776-golang/http"
)

// Comment
func Home(req *http.Request, res *http.Response) *http.Response {
	articles, err := newsapi.TopHeadlines("", "", 50, time.Now().Format("2006-01-02"))

	if err != nil {
		articles = make([]newsapi.Article, 0)
	}

	return res.View("index", http.ViewData{"articles": &articles})
}

// Comment
func Category(req *http.Request, res *http.Response) *http.Response {
	// TODO: Forgot to add params in requests
	return res.View("category", http.ViewData{})
}

// Comment
func Single(req *http.Request, res *http.Response) *http.Response {
	return res.View("single", http.ViewData{})
}
