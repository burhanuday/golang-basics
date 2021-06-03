package main

func main() {
	isTrue := true

	if isTrue == true {
		// do something
	} else if false {
		// do something else
	} else {
		// hereeee
	}

	myVar := "cat"
	// break is not required in go
	// go only executes the first matching case
	switch myVar {
	case "cat":
		// do something
	default:
		// default case
	}
}
