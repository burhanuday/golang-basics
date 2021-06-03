package main

import "fmt"

// interfaces are like contracts
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "SSSS",
		Breed: "Labrador",
	}
	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "Kong",
		Color:         "Brown",
		NumberOfTeeth: 22,
	}
	fmt.Println(gorilla)
	//  throws error
	// PrintInfo(gorilla)
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	fmt.Println("Animal says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}
