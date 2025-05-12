package news

import (
	"fmt"
	newsapi "server/app/services/news_api"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/lucas11776-golang/http"
)

var ArticleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"publisher": &graphql.Field{
				Type: graphql.String,
			},
			"published_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* http://{{ host }}/product?query={articles{publisher,published_at,image,title,description,content,url}} */
			"articles": &graphql.Field{
				Type:        graphql.NewList(ArticleType),
				Description: "Get product list",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					q := ""
					category := ""
					limit := 50

					if s, ok := params.Args["q"].(string); ok {
						q = s
					}

					if c, ok := params.Args["category"].(string); ok {
						category = c
					}

					if l, ok := params.Args["limit"].(int); ok {
						limit = l
					}

					return newsapi.Fetch(q, category, limit), nil
				},
			},
		},
	})

type Article struct {
	Publisher   string    `json:"publisher"`
	PublishedAt time.Time `json:"published_at"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Url         string    `json:"url"`
}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: QueryType,
	},
)

func execute(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

// Comment
func Endpoint(req *http.Request, res *http.Response) *http.Response {
	return res.Json(execute(req.URL.Query().Get("query"), schema))
}
