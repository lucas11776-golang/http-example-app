package bootstrap

import (
	"server/app/controllers"
	"server/env"
	"server/routes"

	"github.com/lucas11776-golang/http"
)

// Comment
func Boot() *http.HTTP {
	server := http.Server(env.Env("HOST"), env.EnvInt("PORT"))

	server.SetStatic(env.Env("ASSETS"))
	server.SetView(env.Env("VIEWS"), env.Env("VIEWS_EXTENSION"))
	server.Session([]byte(env.Env("SESSION_KEY")))

	server.Route().Group("/", routes.Web)
	server.Route().Group("api", routes.Api)
	server.Route().Group("/", routes.Ws)
	server.Route().Fallback(controllers.NotFoundPage)

	return server
}
