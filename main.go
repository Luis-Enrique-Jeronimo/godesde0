package main

import (
	"fmt"

	Variables "github.com/Luis-Enrique-Jeronimo/godesde0/variables"
)

func main() {
	estado, texto := Variables.ConviertoaTexto(140)
	fmt.Println(estado)
	fmt.Println(texto)
}
