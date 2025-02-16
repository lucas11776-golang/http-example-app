package bootstrap

import (
	"server/app/controllers"
	"server/routes"

	"github.com/lucas11776-golang/http"
)

var (
	ADDRESS = "127.0.0.1"
	PORT    = 8080
)

func Boot() *http.HTTP {
	server := http.Server("127.0.0.1", 8080)

	server.SetStatic("static").SetView("views", "html")

	server.Route().Group("/", routes.Web)
	server.Route().Group("api", routes.Api)
	server.Route().Group("/", routes.Ws)
	server.Route().Fallback(controllers.NotFoundPage)

	return server
}
