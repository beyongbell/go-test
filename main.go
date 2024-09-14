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
	// var เปลี่นรค่าได้
	// const เปลี่ยนค่าไม่ได้
	fullname := "Thinnakorn"
	age := 32
	fmt.Printf("Hello %s age %d", fullname, age)
	fullname = "BelL"
	age = age * 2
	fmt.Println(fullname)
	fmt.Println(age)
	fmt.Println()
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)
	sayHello()
	sayHello()
	sayHello()
	bell.SayHelloBell()
	bell.SayHelloBell()
	result := multipleNumber(2, 3)
	fmt.Println("Result:", result)
	sum := addNumbers(5, 7)
	fmt.Println("Sum:", sum)
	for i := 1; i < 10; i++ {
		fmt.Printf("number : %d\n", i)
	}
}
