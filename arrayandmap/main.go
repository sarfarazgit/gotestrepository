package arrayandmap

import "fmt"

func main() {
	// Data prepration for an array
	elements := [10]string{

		"Pen", "Pencil", "Eraser", "Book",
	}
	fmt.Println("\n *** Arrays Elements :*** ")
	for _, val := range elements {
		fmt.Println("\t" + val)
	}

	// Data initialization by using make and display elements
	DisplayElement(ArrayDataInitializationByMake())

	// Map By using make function

	mapelements := make(map[string]string)

	mapelements["Computer"] = "Electronics"
	mapelements["phone"] = "Digital"
	mapelements["Box"] = "Plastics"

	ShowElements(mapelements)

	// Test for element checking
	CheckElements("Test", mapelements)
	CheckElements("Computer", mapelements)

	// Without make function

	newmapelements := map[string]string{
		"pen":    "stationary",
		"Mobile": "Electronics",
	}
	CheckElements("pen", newmapelements)
	CheckElements("Computer", newmapelements)
}
