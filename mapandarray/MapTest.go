package main

import (
	"fmt"
)

func ShowElements(elements map[string]string) {
	fmt.Println("\n **** Map Elements are **** \n")

	for i, val := range elements {
		fmt.Println("["+i+":", val+"]")
	}
}

func checkElements(key string, elements map[string]string) {
	val, isAvailable := elements[key]
	if isAvailable {
		fmt.Println("\nRequested element is available...", key+":"+val)
	} else {
		fmt.Println("\nRequested element is not available...", key+":"+val)
	}
}

func main() {

	// By using make function

	elements := make(map[string]string)

	elements["Computer"] = "Electronics"
	elements["phone"] = "Digital"
	elements["Box"] = "Plastics"

	ShowElements(elements)

	// Test for element checking
	checkElements("Test", elements)
	checkElements("Computer", elements)

	// Without make function

	mapelements := map[string]string{
		"pen":    "stationary",
		"Mobile": "Electronics",
	}
	checkElements("pen", mapelements)
	checkElements("Computer", mapelements)

}
