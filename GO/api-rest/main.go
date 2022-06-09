package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//router := mux.NewRouter().StrictSlash(true)

	/*router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contacto)
	router.HandleFunc("/peliculas", ListaPeliculas)
	router.HandleFunc("/pelicula/{id}", MovieShow)*/

	/*http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(rw, "Hola amo, soy su nuevo servidor en GO...")

	})*/

	router := NewRouter()

	server := http.ListenAndServe(":7777", router)

	log.Fatal(server)

	fmt.Println("El servidor ha sido apagado, amo")
}
