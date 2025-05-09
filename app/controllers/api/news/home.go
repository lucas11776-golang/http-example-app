package news

import (
	newsapi "server/app/services/news_api"

	"github.com/lucas11776-golang/http"
)

// Comment
func Home(req *http.Request, res *http.Response) *http.Response {
	return res.Json(newsapi.Fetch(req.URL.Query().Get("q"), "", 50))
}

// Comment
func Category(req *http.Request, res *http.Response) *http.Response {
	return res.Json(newsapi.Fetch(req.URL.Query().Get("q"), req.Parameters.Get("category"), 50))
}
