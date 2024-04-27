package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println("The value of the pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("The value of the pointer is", ptr)
	fmt.Println("The value inside the pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The New value inside the pointer is", myNumber)

}
