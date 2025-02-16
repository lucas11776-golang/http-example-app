package controllers

import (
	"github.com/lucas11776-golang/http"
)

func NotFoundPage(req *http.Request, res *http.Response) *http.Response {
	return res.View("404", http.ViewData{})
}
