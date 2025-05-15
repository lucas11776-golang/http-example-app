package middlewares

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/types"
)

// Comment
func Cors(req *http.Request, res *http.Response) *http.Response {
	return res.SetHeaders(types.Headers{
		"Access-Control-Request-Method": "GET,POST,DELETE,PUT",
		"Access-Control-Allow-Headers":  "*",
		"Access-Control-Allow-Origin":   "*",
	})
}
