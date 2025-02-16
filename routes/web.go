package routes

import (
	"server/app/controllers/web/home"
	"server/app/middlewares/web"

	"github.com/lucas11776-golang/http"
)

func Web(route *http.Router) {
	route.Get("/", home.Index)

	route.Group("authentication", func(route *http.Router) {
		route.Group("register", func(route *http.Router) {

		}).Middleware(web.IsGuest)

		route.Group("login", func(route *http.Router) {

		}).Middleware(web.IsGuest)

		route.Group("logout", func(route *http.Router) {

		}).Middleware(web.IsUser)
	})
}
