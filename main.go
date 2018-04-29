package main

import (
	"net/http"
	"log"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()

	router.
		Methods("GET").
		MatcherFunc(func(r *http.Request, m *mux.RouteMatch) bool {
		match, err := regexp.MatchString(`/[A-Za-z0-9]+`, r.URL.Path)
		if err != nil {
			panic(err)
		}
		return match
	}).
		Handler(Logger(http.HandlerFunc(RedirectTo), "RedirectTo"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
