package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4{	
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}
	if os.Args[1] == "-h" {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}
	_, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) {
		fmt.Println("Input not found")
		return
	}	
	_, err1 := os.Stat(os.Args[3])
	if os.IsNotExist(err1) {
		fmt.Println("Airport lookup not found")
		return
	}

	airportLookup, err := ReadCsv(os.Args[3])
	if err != nil {
		fmt.Println("Airport lookup malformed")
		return
	}

	malformed := Malformed(airportLookup)
	if malformed {
		fmt.Println("Airport lookup malformed")
		return
	}

	inputString := ReadInput(os.Args[1])
	outputString := changeAirportName(inputString)

	WriteOutput(inputString)

	fmt.Println("ok")
	return
	
}
