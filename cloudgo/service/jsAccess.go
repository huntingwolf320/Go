package service

import (
	"net/http"

	"github.com/unrolled/render"
)

//handle a request with method GET and path "/api/test".
//support js access and return a json data
func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Show string `json:"show"`
		}{Show: "API test!"})
	}
}
