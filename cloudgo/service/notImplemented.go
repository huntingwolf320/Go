package service

import (
	"net/http"
)

func notImplementedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "501 Not Implemented", 501)
	}
}
