package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//                key    value
	myMap := make(map[string]string)
	myMap["dog"] = "Idontlike"
	myMap["cat"] = "Ilike"

	fmt.Println(myMap["dog"])

	//                key    value
	myMapTwo := make(map[string]int)
	myMapTwo["first"] = 1
	myMapTwo["second"] = 2

	fmt.Println(myMapTwo["first"])

	myMapUser := make(map[string]User)

	me := User{
		FirstName: "Burhan",
		LastName:  "Uday",
	}

	myMapUser["me"] = me

	// Slices
	var mySlice []string

	mySlice = append(mySlice, "Hello")
	mySlice = append(mySlice, "World!")

	fmt.Println(mySlice)

	numberSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numberSlice)
	fmt.Println(numberSlice[:2])
}
