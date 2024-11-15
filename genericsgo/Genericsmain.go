package main

import "fmt"

// It supports to (int,string and boolean) type of data:

func displayElements[T int | string | bool](a, b T) {
	fmt.Println("\t displayElements() data \n\t", a, b)
}

// It supports to (any : Whatever it contains the types) type of data:
func arrayElements[T any](elements []T) {

	fmt.Println("\n\t arrayElements() data:\n")

	for _, value := range elements {
		fmt.Print("\t", value, "\n")
	}

}

// To manage Generics by using predefined comparable
func mapElements[T comparable](mapelements map[T]T) {
	fmt.Println("\n\t Map Elements:\n")

	for key, val := range mapelements {
		fmt.Println("\t", key, ":", val, ",")
	}
}

func main() {

	displayElements(false, true)

	// Data preparation for array
	elements := []string{"Hello", "Hi", "Bye"}
	arrayElements(elements)

	// data preparation for map

	mapelements := map[string]string{"1": "Mango", "2": "Banan"}
	mapElements(mapelements)

}
