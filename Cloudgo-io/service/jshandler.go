package service

import (
	"net/http"

	"github.com/unrolled/render"
)

//JsHandler .
func JsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Title      string `json:"title"`
		}{Title: "How are you!"})
	}
}
