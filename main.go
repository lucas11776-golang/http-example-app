package main

import (
	"fmt"
	"server/app/controllers"
	"server/routes"

	"github.com/lucas11776-golang/http"
)

func main() {
	server := http.Server("127.0.0.1", 8080)

	server.SetStatic("static").SetView("views", "html")

	server.Route().Group("/", routes.Web)
	server.Route().Group("api", routes.Api)
	server.Route().Group("/", routes.Ws)
	server.Route().Fallback(controllers.NotFoundPage)

	fmt.Printf("Server Running %s", server.Host())

	server.Listen()
}
