package main

import (
	"fmt"

	"github.com/beyongbell/go-test/bell"
	"github.com/google/uuid"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func multipleNumber(a int, b int) int {
	return a * b
}

func addNumbers(a, b int) int {
	return a + b
}

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)
	sayHello()
	sayHello()
	sayHello()
	bell.SayHelloBell()
	result := multipleNumber(2, 3)
	fmt.Println("Result:", result)
	sum := addNumbers(5, 7)
	fmt.Println("Sum:", sum)
}
