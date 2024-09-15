package main

import (
	"fmt"

	"github.com/beyongbell/go-test/bell"
	"github.com/google/uuid"
)

type Person struct {
	Name    string
	Age     int
	Address any
}

type Address struct {
	Street  string
	City    string
	Zipcode int
}

type Student struct {
	Firstname string
	Lastname  string
}

func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func multipleNumber(a int, b int) int {
	return a * b
}

func addNumbers(a, b int) int {
	return a + b
}

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type PersonS struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p PersonS) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

// divide divides two integers and returns an error if the divisor is 0
// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

type LoginError struct {
	Username string
	Message  string
}

// Implement the Error() method to satisfy the error interface
func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

// Simulated function that attempts a user login
func login(username, password string) error {
	// Simulate checking username and password
	if username != "admin" || password != "password123" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	// Login successful
	return nil
}

func main() {
	var person Person
	person.Name = "thinnakorn"
	person.Age = 32
	person.Address = Address{
		Street:  "21 m2",
		City:    "ss",
		Zipcode: 12345,
	}

	bob := Person{
		Name: "Bob",
		Age:  22,
		Address: Address{
			Street:  "21 m2",
			City:    "ss",
			Zipcode: 12345,
		},
	}
	fmt.Println(person)
	fmt.Println(bob)
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
	sayHello("Bell")
	sayHello("Jonh")
	sayHello("Dev")
	bell.SayHelloBell()
	bell.SayHelloBell()
	result := multipleNumber(2, 3)
	fmt.Println("Result:", result)
	sum := addNumbers(5, 7)
	fmt.Println("Sum:", sum)
	for i := 1; i < 10; i++ {
		fmt.Printf("number : %d\n", i)
	}

	var myArray [3]string
	myArray[0] = "bell"
	myArray[1] = "thinnakorn"
	myArray[2] = "joompee"
	fmt.Println(myArray)

	student := Student{
		Firstname: "BelL",
		Lastname:  "NaJA",
	}

	// Call the FullName method on the Student instance
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)

	dog := Dog{Name: "Buddy"}
	persons := PersonS{Name: "Alice"}

	makeSound(dog)
	makeSound(persons)

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Result:", result)

	// Attempt to login with wrong credentials
	err := login("user", "pass")
	if err != nil {
		switch e := err.(type) {
		case *LoginError:
			// Custom error handling
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)
		}
		return
	}

	// Continue with the rest of the program if login is successful
	fmt.Println("Login successful!")
}
