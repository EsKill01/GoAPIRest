package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inicio del programa")

	fmt.Print("Pedido 1 -->")
	fmt.Println(gorras(52, "$"))

	fmt.Print("Pedido 2 -->")
	fmt.Println(gorras(100, "$"))

	fmt.Print("Pedido 3 -->")
	fmt.Println(gorras(98, "$"))

	pantalon("rojo", "largo", "sin bolsillos", "nike")

	var pelicula [3]string
	pelicula[0] = "Batman"
	pelicula[1] = "La liga de la justicia"
	pelicula[2] = "Joker"

	peliculas := [3]string{"La verdad duele", "La leyenda de los dragones", "Alibaba y los 40 ladrones"}

	var lasPeliculas [3][2]string

	lasPeliculas2 := []string{
		"Siempre arriba la ultima batalla",
		"Por siempre tuyo",
		"El ostion de tierra",
		"Fenix"}

	fmt.Println("Lista de objetos-->")

	for _, v := range lasPeliculas2 {
		fmt.Println(v)
	}

	imprimirMensaje("Se agregar mas a la lista")

	lasPeliculas2 = append(lasPeliculas2, "Alimentados en las tinieblas")
	lasPeliculas2 = append(lasPeliculas2, "Alimentados en las tinieblas2")
	lasPeliculas2 = append(lasPeliculas2, "Alimentados en las tinieblas3")

	fmt.Println(len(lasPeliculas))
	fmt.Println(len(lasPeliculas[0:3]))

	fmt.Println(lasPeliculas2)

	imprimirMensaje("")

	lasPeliculas[0][0] = "Siempre arriba"
	lasPeliculas[0][1] = "Los colonos artistas"
	lasPeliculas[1][0] = "Siempre arriba 2"
	lasPeliculas[1][1] = "La verguenza de los caidos"
	lasPeliculas[2][0] = "Mintiendome"
	lasPeliculas[2][1] = "La respuesta acertiva"

	fmt.Println("Libros array dimencional")

	fmt.Println("--------------------------------")

	count := 0

	for i := 0; i < len(lasPeliculas); i++ {

		fmt.Println("El item del array en la primer posición " + lasPeliculas[i][count])

		for y := 0; y < len(lasPeliculas[i]); y++ {

			if i > 0 || y > 0 {

				fmt.Println("El item del array en la segunda posicion " + lasPeliculas[i][y])

				if count == 0 {
					count++
				}
			}

		}

	}

	/*for _, item := range lasPeliculas {

		fmt.Println("El item del array en la primer posición " + string(count) + " " + item[count])

		for _, item2 := range item[count] {

			fmt.Println("El item del array en la segunda posicion " + string(item2))

		}
	}*/

	fmt.Println("--------------------------------")

	for _, item := range peliculas {
		fmt.Println(item)
	}

	for _, item := range peliculas {
		fmt.Println(item)
	}

	fmt.Println(pelicula[1])
	fmt.Println(peliculas[2])

}

func pantalon(caracteristicas ...string) {

	for _, item := range caracteristicas {
		fmt.Println(item)
	}

}

func imprimirMensaje(mensaje string) {
	fmt.Println(mensaje)
}

func gorras(pedido int, moneda string) (string, int, string) {

	precio := func() float32 {
		return float32(pedido) * 7
	}

	return "El precio del pedido es: ", int(precio()), moneda
}
