package ejercicios

import (
	"strconv"
)

func Publica(cadena string) (int, string) {

	var numero int
	numero, _ = strconv.Atoi(cadena)

	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
