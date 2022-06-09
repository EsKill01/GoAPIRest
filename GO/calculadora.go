package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inicio del programa")

	var num1 float32 = 10
	var num2 float32 = 6

	calculadora(num1, num2)
	calculadora(458, 100)
	calculadora(47, 52)
	calculadora(27, 91)
	holaMundo()
	fmt.Println()
	fmt.Println(devolverText())
	fmt.Println(devolverText2())
}

func holaMundo() {
	fmt.Println("Hola calculadora")
}

func devolverText() (string, string) {

	dato1 := "Esteba"
	dato2 := "Barboza"
	return dato1, dato2
}

func devolverText2() (dato1 string, dato2 string) {

	dato1 = "Jesa"
	dato2 = "Martinez"
	return
}

func operacion(num1 float32, num2 float32, op string) float32 {
	var resultado float32

	if op == "+" {
		resultado = num1 + num2
	}

	switch op {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "/":
		resultado = num1 / num2
	case "*":
		resultado = num1 * num2

	}

	return resultado
}

func calculadora(num1 float32, num2 float32) {

	fmt.Print("La suma es: ")
	fmt.Println(operacion(num1, num2, "+"))

	fmt.Print("La resta es: ")
	fmt.Println(operacion(num1, num2, "-"))

	fmt.Print("La multiplicación es: ")
	fmt.Println(operacion(num1, num2, "*"))

	fmt.Print("La division es: ")
	fmt.Println(operacion(num1, num2, "/"))

	fmt.Print("el resto es: ")
	fmt.Println(int(num1) % int(num2))

}
