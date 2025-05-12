package routes

import (
	// "server/app/controllers/api/news"

	"server/app/controllers/graph_ql/news"

	"github.com/lucas11776-golang/http"
)

// Comment
func GraphQL(route *http.Router) {
	route.Get("news", news.Endpoint)
}
