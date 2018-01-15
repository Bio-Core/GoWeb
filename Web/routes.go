package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PersonList",
		"GET",
		"/people",
		PersonList,
	},
	Route{
		"PersonListJSON",
		"GET",
		"/peopleJSON",
		PersonListJSON,
	},
	Route{
		"PersonCreate",
		"POST",
		"/people",
		PersonCreate,
	},
	Route{
		"GenericJSON",
		"POST",
		"/JSON",
		GenericJSON,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/JSON",
		GenericListJSON,
	},
}