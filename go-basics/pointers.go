package main

import "fmt"

func main() {
	myString := "Burhan"
	changeUsingPointer(&myString)
	fmt.Println(myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
