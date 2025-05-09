package routes

import (
	"server/app/controllers/web/news"

	"github.com/lucas11776-golang/http"
)

// Comment
func Web(route *http.Router) {
	route.Get("/", news.Home)
	route.Group("/", func(route *http.Router) {
		route.Group("categories", func(route *http.Router) {
			route.Get("{category}", news.Category)
		})
	})
}
