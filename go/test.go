package main

import (
	"fmt"
	"time"
)

// whatAmI is a function that takes one argument i of type interface{}.
// The empty interface can hold any type of value.
func whatAmI(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

// Interface and Struct Definitions
type Animal interface {
	Speak() string
}

// This line defines a new type named Dog.
// A struct is a composite data type that groups together variables under a single name for creating more complex data structures.
// struct{} denotes that Dog is a struct type with no fields.
type Dog struct{}

// func (d Dog) specifies the method receiver. In this case, d is the receiver of type Dog. This means Speak is a method that can be called on any value of type Dog.
func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	fmt.Println("hello world")

	// LOOP: FOR
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println(time.Sunday + 1)
	for i := time.Sunday; i <= time.Saturday; i++ {
		fmt.Println(i)
	}

	// IF/ELSE
	i = 0
	if i == 0 {
		fmt.Println(i)
	}

	// SWITCH
	i = 2

	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch t.Year() {
	case 2025:
		fmt.Print(t)
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// Creating an instance of Dog and calling Speak method
	var myDog Dog
	fmt.Println(myDog.Speak()) // Output: Woof!
}
