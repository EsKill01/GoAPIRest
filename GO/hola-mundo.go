package main

import (
	"fmt"
	//"time"
)

func main() {

	/*Comentarios de varias lineas

	sss

	*/

	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Esteban"

	var apellidos string = "Barboza"
	apellidos = "Barboza Muñoz"

	//Costantes segun el tipo
	pais := "Costa Rica"

	pais = "España"

	var prueba float32 = 2.3

	const year int = 2022

	fmt.Println("Voy a dormirme un rato")
	//time.Sleep(time.Second * 5)
	fmt.Println("Hola mundo")
	fmt.Println(year)

	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println("Hola mi dueño: " + nombre + " " + apellidos + "" + pais)
	fmt.Println(prueba)

}
