package middleware

import "github.com/lucas11776-golang/http"

// Comment
func IsGuest(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}

// Comment
func IsAuth(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}
