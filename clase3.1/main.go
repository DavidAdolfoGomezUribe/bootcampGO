package main

import (
    "fmt"
    "example.com/fizzbuzz/fizzbuzz"
)

func main() {
    for i := 1; i <= 20; i++ {
        fmt.Println(fizzbuzz.FizzBuzz(i))
    }
}
