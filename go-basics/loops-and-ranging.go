package main

import (
	"fmt"
)

func main() {
	// go only has for loop
	count := 5
	for i := 0; i < count; i++ {
		// do something count times
	}

	mySlice := []string{"cat", "dog", "fish", "banana"}
	for i, item := range mySlice {
		fmt.Println(i, item)
	}

	myMap := make(map[string]string)

	myMap["animal"] = "Dog"
	for _, item := range myMap {
		fmt.Println(item)
	}

}
