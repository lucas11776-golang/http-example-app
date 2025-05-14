package routes

import (
	// "server/app/controllers/api/news"

	"server/app/controllers/graph_ql/news"
	graphql "server/utils/graph_ql"

	"github.com/lucas11776-golang/http"
)

// Comment
func GraphQL(route *http.Router) {
	route.Post("news", graphql.GraphQLRoute(news.Home()))
}
