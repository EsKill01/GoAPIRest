package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("-----Inicio del programa")

	nuevoTexto := os.Args[1] + "\n"

	fmt.Println("-----El amo ingreso el nuevo texto: " + string(nuevoTexto))

	archivo, err := os.OpenFile("archivo.txt", os.O_APPEND, os.ModeAppend)
	escribirOk, err := archivo.WriteString(string(nuevoTexto))
	archivo.Close()

	fmt.Println(string(escribirOk))

	showError(err)

	//escribir := ioutil.WriteFile("archivoNuevo.txt", nuevoTexto, os.ModeAppend)

	//showError(escribir)

	fmt.Println("Listo amo, ya se agrego el texto")

	text, error := ioutil.ReadFile("archivo.txt")
	showError(error)

	fmt.Println(string(text))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
