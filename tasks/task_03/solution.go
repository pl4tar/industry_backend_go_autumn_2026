package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}

	n3 := n % 3
	n5 := n % 5

	switch {
	case n3 == 0 && n5 == 0:
		return "FizzBuzz", nil
	case n3 == 0:
		return "Fizz", nil
	case n5 == 0:
		return "Buzz", nil
	default:
		return strconv.Itoa(n), nil
	}
}
