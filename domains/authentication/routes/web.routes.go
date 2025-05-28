package routes

import (
	"server/domains/authentication/controller/web/login"
	"server/domains/authentication/controller/web/logout"
	"server/domains/authentication/controller/web/register"
	"server/domains/authentication/middleware"

	"github.com/lucas11776-golang/http"
)

// Comment
func Web(route *http.Router) {
	route.Group("login", func(route *http.Router) {
		route.Get("/", login.View)
		route.Post("/", login.Store)
	}, middleware.IsGuest)

	route.Group("register", func(route *http.Router) {
		route.Get("/", register.View)
		route.Post("/", register.Store)
	}, middleware.IsGuest)

	route.Group("logout", func(route *http.Router) {
		route.Delete("/", logout.Destroy)
	}, middleware.IsAuth)
}
