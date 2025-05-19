package news

import (
	"math/rand"
	newsapi "server/app/services/news_api"

	"github.com/graphql-go/graphql"
)

type Rating struct {
	Score int `json:"score"`
	Votes int `json:"votes"`
}

// Define Rating type
var ratingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Rating",
	Fields: graphql.Fields{
		"score": &graphql.Field{
			Type: graphql.Int,
		},
		"votes": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"rating": &graphql.Field{
				Type: ratingType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Will call service
					rating := Rating{
						Score: rand.Intn(10) % 100,
						Votes: rand.Intn(1000000),
					}
					return rating, nil
				},
			},

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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"articles": &graphql.Field{
				Type:        graphql.NewList(articleType),
				Description: "Get product list",
				Args: graphql.FieldConfigArgument{
					"q": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					q := ""
					category := ""
					limit := 100

					if s, ok := params.Args["q"].(string); ok {
						q = s
					}

					if c, ok := params.Args["category"].(string); ok {
						category = c
					}

					if l, ok := params.Args["limit"].(int); ok {
						limit = l
					}

					return newsapi.FetchHeadlinesLatest(q, category, limit), nil
				},
			},
		},
	})

// Comment
func Home() graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	return schema
}
