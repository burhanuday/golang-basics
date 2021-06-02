package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	result := sayHello("Burhan")

	fmt.Println((result))
}

func sayHello(s string) string {
	fmt.Println(s + " says hello")

	return s + " says hello"
}
