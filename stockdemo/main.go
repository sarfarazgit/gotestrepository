package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("\n ***  Welcome To The Stock Exchange Marcket ***")

	// Preparing the Stock data

	stockdata := Stock{"SBI stock", 150}

	// Sending for conversion (Marshaling)
	var jsondata = StructureTojson(stockdata)

	fmt.Printf("Your Stock details in jsondata: %v\n", jsondata)

	// Sending for conversion (Unmarshalling)
	JsonTostructure([]byte(jsondata))
}

// Converting structure data to json formate (Marshaling)

func StructureTojson(stockdata Stock) string {

	var jsonresult, err = json.Marshal(stockdata)

	if err != nil {

		return ("Invalid data recieved for conversion:")
	}

	return string(jsonresult)
}

// Converting JSON data into Structure formate (Unmarshaling)

func JsonTostructure(Jsondata []byte) {

	var stockdata Stock
	erorr := json.Unmarshal(Jsondata, &stockdata)

	if erorr != nil {

		fmt.Println("Invalid data recieved for conversion:", erorr)
		return
	}

	fmt.Println("Your stock data post conversion:", stockdata)
}

type Stock struct {
	Stockname  string
	TotalStock int
}
