package main

import (
	"fmt"
	"strconv"
)

var val = "Global variable"

func printGlobal() {
	fmt.Println("Val: ", val)
}

func ArraySkice() {
	// Array size will increase automatically based on the requirement (Dynamic array)
	names := make([]string, 0)

	names = append(names, "sarfaraz")
	names = append(names, "Sudhish")
	names = append(names, "Anand")

	fmt.Println(names)

	// Slices

	fmt.Println(names[:3])
	fmt.Println(names[:3])
	fmt.Println(names[:2])
}

func main() {

	// typeconversion()
	// printGlobal()

	// fmt.Println("Current Date is....", servicepackage.CurrentDate())
	// fmt.Println("Current Time is....", servicepackage.CurrenTime())

	// auther := Auther{
	// 	Name:    "John1",
	// 	Address: "USA",
	// }

	// fmt.Println(auther.GetAutherDetails())
	// auther.SetAutherDetails("John2", "USA")
	// fmt.Println(auther.GetAutherDetails())

	// ArraySkice()
}

type Auther struct {
	Name, Address string
}

func (auther Auther) GetAutherDetails() string {
	return "Auther name is:" + auther.Name + "\t Auther address is:" + auther.Address
}

func (auther *Auther) SetAutherDetails(Name, Address string) {
	auther.Name = Name
	auther.Address = Address
}

func typeconversion() {
	a := 100
	b := uint(a)

	c := byte(b)

	fmt.Println("b = ", b)
	fmt.Println("byte val c = ", c)

	d := float64(c)
	fmt.Println("Float val d = ", d)

	msg := "2000"
	intval, _ := strconv.Atoi(msg)
	fmt.Println("After conversion...", intval)

}
