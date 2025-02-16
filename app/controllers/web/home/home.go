package home

import (
	"github.com/lucas11776-golang/http"
)

func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("home", http.ViewData{})
}
