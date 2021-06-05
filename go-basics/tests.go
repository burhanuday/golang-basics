package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100, 10)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result", result)
}

func divide(x float32, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}

	result = x / y

	return result, nil
}
