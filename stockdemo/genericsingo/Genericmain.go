package main

import "fmt"

func displayElements[T int | string | bool](a, b T) {
	fmt.Println(a, b)
}

func arrayElements[T any](elements []T) {

	for val := range elements {
		fmt.Println(val)
	}

}

func mapElements[T comparable](mapelements map[T]T) {
	for key, val := range mapelements {
		fmt.Println(key, val)
	}
}

func main() {

	displayElements(100, 200)

	// Data preparation for array
	elements := []string{"Hello", "Hi", "Bye"}
	arrayElements(elements)

	// data preparation for map

	mapelements := map[string]string{"1": "Mango", "2": "Banan"}
	mapElements(mapelements)

}
