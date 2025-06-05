package routes

import (
	"server/app/controllers/api/news"
	"server/domains/authentication/routes"

	"github.com/lucas11776-golang/http"
)

// Comment
func Api(route *http.Router) {
	route.Group("news", func(route *http.Router) {
		route.Get("/", news.Home)
		route.Group("categories", func(route *http.Router) {
			route.Get("{category}", news.Category)
		})
	})
}

// Comment
func API(route *http.Router) {
	route.Group("api", func(route *http.Router) {
		route.Group("authentication", routes.API)
	})
}
