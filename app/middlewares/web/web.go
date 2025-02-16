package web

import (
	"github.com/lucas11776-golang/http"
)

func IsGuest(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}

func IsUser(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}
