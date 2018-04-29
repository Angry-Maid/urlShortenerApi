package main

import "net/http"

type Route struct {
	Name 		string
	Method 		string
	Pattern	   	string
	HandleFunc 	http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GenerateShortURL",
		"POST",
		"/generate",
		GenerateShortURL,
	},
}
