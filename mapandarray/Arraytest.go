package main

import "fmt"

func main() {
	// Data prepration for an array
	elements := [10]string{

		"Pen", "Pencil", "Eraser", "Book",
	}
	fmt.Println("\n *** Arrays Elements :*** \n")
	for _, val := range elements {
		fmt.Println("\t" + val)
	}
	// Function calling
	UsingMake()
}

func UsingMake() {
	element := make([]int, 5)
	element[0] = 100
	element[1] = 200
	element[2] = 300
	element[3] = 400
	element[4] = 500
	// Display all elements
	fmt.Println("\n *** Arrays Elements By using make:***")
	for _, value := range element {
		fmt.Println(value)
	}
}
