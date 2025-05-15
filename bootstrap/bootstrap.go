package bootstrap

import (
	"server/app/controllers"
	"server/app/middlewares"
	"server/app/middlewares/api"
	"server/env"
	"server/routes"

	"github.com/lucas11776-golang/http"
)

// Comment
func Boot() *http.HTTP {
	server := http.Server(env.Env("HOST"), env.EnvInt("PORT"))

	configureOptions(server)
	configureRoutes(server)

	return server
}

// Comment
func configureOptions(server *http.HTTP) {
	server.SetStatic(env.Env("ASSETS"))                          // Static Assets
	server.SetView(env.Env("VIEWS"), env.Env("VIEWS_EXTENSION")) // Views
	server.Session([]byte(env.Env("SESSION_KEY")))               // Session
}

// Comment
func configureRoutes(server *http.HTTP) {
	server.Route().Options("*", middlewares.Cors) // Preflight
	server.Route().Group("/", routes.Web)         // Web
	server.Route().Group("/", routes.Ws)          // Websocket
	server.Route().Middleware(api.CorsMiddleware).Group("/", func(route *http.Router) {
		server.Route().Group("api", routes.Api)          // API
		server.Route().Group("graph_ql", routes.GraphQL) // GraphQL
	})
	server.Route().Fallback(controllers.NotFoundPage) // Page Not Found
}
