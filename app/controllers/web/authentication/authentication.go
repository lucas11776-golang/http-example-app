package authentication

import "github.com/lucas11776-golang/http"

func LoginPage(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.login", http.ViewData{})
}

func RegisterPage(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.register", http.ViewData{})
}
