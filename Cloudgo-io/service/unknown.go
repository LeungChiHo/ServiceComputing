package service

import (
	"net/http"

	"github.com/unrolled/render"
)

//UnknownHandler .
func UnknownHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "501 Not Implemented", 501)
	}
}
