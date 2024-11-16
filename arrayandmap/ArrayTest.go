package arrayandmap

import (
	"fmt"
	"math/rand/v2"
)

func ArrayDataInitializationByMake() []int {
	elements := make([]int, 5)
	for i := 0; i < len(elements); i++ {
		elements[i] = rand.IntN(50)
	}
	return elements
}

func DisplayElement(elements []int) {
	// Display all elements
	fmt.Println("\n *** Arrays Elements By using make:***")
	for _, value := range elements {
		fmt.Println(value)
	}
}
