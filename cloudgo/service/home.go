package service

import (
	"net/http"

	"text/template"
)

//handle a request with method GET and path "/".
//return homepage
func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		t.Execute(w, nil)
	}
}

//handle a request with method POST and path "/".
//add the input data to the homepage
func postHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		t := template.Must(template.ParseFiles("templates/index.html"))
		t.Execute(w, map[string]string{
			"name": req.Form["name"][0],
			"word": req.Form["word"][0],
		})
	}
}
