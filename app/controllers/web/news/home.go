package news

import (
	newsapi "server/app/services/news_api"
	"strings"

	"github.com/lucas11776-golang/http"
)

// Comment
func Home(req *http.Request, res *http.Response) *http.Response {
	return res.View("index", http.ViewData{
		"articles": newsapi.FetchHeadlinesLatest(req.URL.Query().Get("q"), "", 100),
		"search":   req.URL.Query().Get("q"),
	})
}

// Comment
func Category(req *http.Request, res *http.Response) *http.Response {
	return res.View("category", http.ViewData{
		"category": strings.Title(req.Parameters.Get("category")),
		"articles": newsapi.FetchHeadlinesLatest(req.URL.Query().Get("q"), req.Parameters.Get("category"), 100),
		"search":   req.URL.Query().Get("q"),
	})
}
