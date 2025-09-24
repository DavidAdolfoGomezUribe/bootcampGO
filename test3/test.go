package main

import ( 
	"fmt"
	"os"
	"strings"
)

func main(){

	var palabra string = strings.Join(os.Args[1:], " ")
	var otrapalabras = []rune(palabra)
	var arbalap []string

	for i := (len(palabra)-1); i >= 0 ; i-- {
		
		//fmt.Println(string(otrapalabras[i]))
		arbalap = append(arbalap, string(otrapalabras[i]))
	}
	
	fmt.Println(strings.Join(arbalap,""))
}


// funcion main de pruebas
// func main(){

// 	fmt.Println("ok")
// 	for i := 10; i >= 1; i-- {
// 		fmt.Println(i)
// 	}
	
// 	var palabra string = strings.Join(os.Args[1:], " ")

// 	palabras := []string{palabra}
// 	var otrapalabras = []rune(palabra)

// 	fmt.Println((palabras[0]))
// 	fmt.Println((otrapalabras[0]))
// 	fmt.Println(string(otrapalabras[0]))

	
// }