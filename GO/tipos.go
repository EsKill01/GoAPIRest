package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	fmt.Println("Inicio del programa")

	var gorraNegra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 22.59,
		plana:  false}

	fmt.Println(gorraNegra)
	fmt.Println(gorraNegra.marca)
}
