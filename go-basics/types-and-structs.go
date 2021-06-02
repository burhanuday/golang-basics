package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (user *User) printFirstName() string {
	return user.FirstName
}

func main() {
	user := User{
		FirstName: "Burhanuddin",
		LastName:  "Udaipurwala",
	}

	fmt.Println(user.FirstName)
	fmt.Println(user.printFirstName())
}
