package service

import (
	"net/http"

	"github.com/unrolled/render"
)

// home .
func home(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		formatter.HTML(w, http.StatusOK, "table", struct {
			U    string
			P    string
		}{U: req.Form["username"][0], P: req.Form["password"][0]})
	}
}
