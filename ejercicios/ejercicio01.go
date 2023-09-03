package ejercicios

import (
	"fmt"
	"strconv"
)

func Publica(cadena string) (int, string) {

	var numero int
	numero, _ = strconv.Atoi(cadena)

	if numero > 100 {
		fmt.Println("Es mayor a 100")
	} else {
		fmt.Println("Es menor a 100")
	}

	return numero, "Hola Mundo!"
}
