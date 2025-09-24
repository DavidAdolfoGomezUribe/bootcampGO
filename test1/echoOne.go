//Esto devuelve los argumentos que se introduzcan en consola 
//go run test1/echoOne.go esto es una prueba --> esto es una prueba


package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separacion string
	for i := 1; i < len(os.Args); i++ {
		s += separacion + os.Args[i]
		separacion = "patata"
	}
	fmt.Println(s)
}
