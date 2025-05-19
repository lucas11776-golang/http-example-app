package bootstrap

import (
	"server/app/controllers"
	"server/app/middlewares"
	"server/app/middlewares/api"
	"server/database"
	"server/env"
	"server/routes"

	"github.com/lucas11776-golang/http"
)

// Comment
func Boot(envPath string) *http.HTTP {
	env.Load(envPath)

	server := http.Server(env.Env("HOST"), env.EnvInt("PORT"))

	configureOptions(server)
	configureDatabase()
	configureRoutes(server)

	return server
}

// Comment
func configureOptions(server *http.HTTP) {
	server.SetStatic(env.Env("ASSETS"))                          // Static Assets
	server.SetView(env.Env("VIEWS"), env.Env("VIEWS_EXTENSION")) // Views
	server.Session([]byte(env.Env("SESSION_KEY")))               // Session
}

func configureDatabase() {
	database.Setup()
}

// Comment
func configureRoutes(server *http.HTTP) {
	server.Route().Options("*", middlewares.Cors) // Preflight
	server.Route().Group("/", routes.Web)         // Web
	server.Route().Group("/", routes.Ws)          // Websocket
	server.Route().Middleware(api.CorsMiddleware).Group("/", func(route *http.Router) {
		route.Group("api", routes.Api)     // API
		route.Group("gql", routes.GraphQL) // GraphQL
	})
	server.Route().Fallback(controllers.NotFoundPage) // Page Not Found
}
