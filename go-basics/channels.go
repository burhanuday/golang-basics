package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numPool = 10

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)

	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	fmt.Println(num)
}
