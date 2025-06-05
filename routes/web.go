package routes

import (
	"server/app/controllers/web/news"

	"github.com/lucas11776-golang/http"
)

// Comment
func Web(route *http.Router) {
	route.Get("/", news.Home)
	route.Get("graph_ql", func(req *http.Request, res *http.Response) *http.Response {
		return res.View("graph_ql", http.ViewData{})
	})
	route.Group("categories", func(route *http.Router) {
		route.Get("{category}", news.Category)
	})
}

func WEB(route *http.Router) {

}
