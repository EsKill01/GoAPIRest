package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Inicio del programa")

	numero, _ := strconv.Atoi(os.Args[3])

	if numero%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	fmt.Println(os.Args)

	fmt.Println("Hola amo: " + os.Args[1] + ". bienvenido a sus apocentos virtuales...")

	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad <= 99 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor  de edad")
	}

	peliculas := []string{"Horror en la granaja", "a trabajar culeros", "Mi novia se apaciona"}

	for i := 0; i < len(peliculas); i++ {
		fmt.Println(peliculas[i])
	}

	for _, item := range peliculas {
		fmt.Println(item)
	}

	momento := time.Now().Weekday()

	switch momento {
	case time.Sunday:
		fmt.Println("Hoy es domingo")
	case time.Monday:
		fmt.Println("Hoy es lunes")

	case time.Tuesday:
		fmt.Println("Hoy es martes")
	case time.Wednesday:
		fmt.Println("Hoy es miercoles")
	case time.Thursday:
		fmt.Println("Hoy es jueves")
	case time.Friday:
		fmt.Println("Hoy es viernes")
	case time.Saturday:
		fmt.Println("Hoy es sabado")
	default:
		fmt.Println("Es otro dia de la semana")

	}

}
