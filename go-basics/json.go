package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Burhan"
		}
	]
	`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
	}
	fmt.Println(unmarshalled)

	// struct to json
	var mySlice []Person

	m1 := Person{
		FirstName: "AAAA",
	}

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {

	}

	fmt.Println(string(newJson))
}
