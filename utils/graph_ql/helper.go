package graphql

import (
	"encoding/json"
	"io"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/lucas11776-golang/http"
)

type Query struct {
	Query string `json:"query"`
}

// Comment
func GraphQLRoute(schema graphql.Schema) http.WebCallback {
	return func(req *http.Request, res *http.Response) *http.Response {
		query, err := GraphQLQuery(req.Body)

		if err != nil {
			return res.SetStatus(http.HTTP_RESPONSE_UNPROCESSABLE_CONTENT).Json(map[string]string{
				"message": "invild query scheme",
			})
		}

		return res.Json(graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query.Query,
		}))
	}
}

// TODO: Postman, EchoAPI
// Comment
func GraphQLQuery(body io.Reader) (*Query, error) {
	data, err := io.ReadAll(body)

	if err != nil {
		return nil, err
	}

	var query Query

	err = json.Unmarshal(data, &query)

	if err != nil {
		return nil, err
	}

	return &query, nil
}

// Comment
func GraphQLErrorResponse(errors ...error) graphql.Result {
	result := graphql.Result{}

	for _, err := range errors {
		result.Errors = append(result.Errors, gqlerrors.FormattedError{
			Message: err.Error(),
		})
	}

	return result
}
