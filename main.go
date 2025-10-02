package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "-h" {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	if len(os.Args) != 4{
		fmt.Println("error")
		return
	}
	fmt.Println("worked")
}
