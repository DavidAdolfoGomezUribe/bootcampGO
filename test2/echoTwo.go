package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Une todos los argumentos desde os.Args[1:] con espacios
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[0:], " "))

}
