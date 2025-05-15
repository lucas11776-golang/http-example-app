package api

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/types"
)

// Comment
func CorsMiddleware(req *http.Request, res *http.Response, next http.Next) *http.Response {
	res.SetHeaders(types.Headers{
		"Access-Control-Allow-Origin":   "*",
		"Access-Control-Allow-Headers":  "*",
		"Access-Control-Request-Method": "GET,POST,DELETE,PUT",
	})

	return next()
}
