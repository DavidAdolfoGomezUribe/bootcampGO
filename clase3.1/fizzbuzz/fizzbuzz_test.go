package fizzbuzz

import "testing"

func TestFizzBuzz_Basicos(t *testing.T) {
    testCases := []struct {
        name     string
        in       int
        expected string
    }{
        {"1 -> '1'", 1, "1"},
        {"2 -> '2'", 2, "2"},
        {"3 -> 'Fizz'", 3, "Fizz"},
        {"5 -> 'Buzz'", 5, "Buzz"},
        {"15 -> 'Fizz Buzz'", 15, "Fizz Buzz"},
        {"30 -> 'Fizz Buzz'", 30, "Fizz Buzz"},
        {"4 -> '4'", 4, "4"},
    }

    for _, tc := range testCases {
        tc := tc // captura de rango
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            got := FizzBuzz(tc.in)
            if got != tc.expected {
                t.Fatalf("FizzBuzz(%d) = %q; want %q", tc.in, got, tc.expected)
            }
        })
    }
}
