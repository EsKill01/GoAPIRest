package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router

}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Contacto",
		"GET",
		"/contact",
		Contacto,
	},
	Route{
		"Peliculas",
		"GET",
		"/peliculas",
		ListaPeliculas,
	},
	Route{
		"Pelicula",
		"GET",
		"/pelicula/{id}",
		MovieShow,
	},
	Route{
		"AddMovie",
		"POST",
		"/pelicula",
		AddMovie,
	},
	Route{
		"UpdateMovie",
		"PUT",
		"/pelicula/{id}",
		UpdateMovie,
	},
	Route{
		"DeleteMovie",
		"DELETE",
		"/pelicula/{id}",
		DeleteMovie,
	},
}
