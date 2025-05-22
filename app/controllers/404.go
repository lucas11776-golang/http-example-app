package controllers

import (
	"github.com/lucas11776-golang/http"
)

// Comment
func NotFoundPage(req *http.Request, res *http.Response) *http.Response {
	if req.GetHeader("Content-Type") == "application/json" {
		return res.SetStatus(http.HTTP_RESPONSE_NOT_FOUND).
			Json(map[string]string{"message": "Route not found"})
	}

	return res.View("404", http.ViewData{"search": req.URL.Query().Get("q")})
}
