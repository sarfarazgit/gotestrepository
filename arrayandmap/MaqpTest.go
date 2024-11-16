package arrayandmap

import (
	"fmt"
)

func ShowElements(elements map[string]string) {
	fmt.Println("\n **** Map Elements are **** ")

	for i, val := range elements {
		fmt.Println("["+i+":", val+"]")
	}
}

func CheckElements(key string, elements map[string]string) {
	val, isAvailable := elements[key]
	if isAvailable {
		fmt.Println("\nRequested element is available...", key+":"+val)
	} else {
		fmt.Println("\nRequested element is not available...", key+":"+val)
	}
}
