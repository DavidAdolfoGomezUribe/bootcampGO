package fizzbuzz

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFizzBuzz_Basicos(t *testing.T) {
    tests := []struct {
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

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            out := FizzBuzz(tc.in)
            assert.Equal(t, tc.expected, out)
        })
    }
}
