package fizzbuzz

import "strconv"

// FizzBuzz aplica las reglas:
// - "Fizz Buzz" si n es múltiplo de 3 y 5
// - "Fizz" si n es múltiplo de 3
// - "Buzz" si n es múltiplo de 5
// - n como string en otro caso
func FizzBuzz(n int) string {
    switch {
    case n%15 == 0:
        return "Fizz Buzz"
    case n%3 == 0:
        return "Fizz"
    case n%5 == 0:
        return "Buzz"
    default:
        return strconv.Itoa(n)
    }
}
