package main

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lookup, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lookup, nil
}

func Malformed(airportLookup [][]string) (bool) {
	if len(airportLookup) == 0 {
		return true
	}

	exampleRow := []string {"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}
	
	i := 0
	for _, row := range airportLookup {
 		if len(row) != 6 {
			return true
		}
		
		for _, column := range row {
			j := 0
			if column == ""{
				return true
			}
			if i == 0 {
				if row[j] != exampleRow[j]{
					return true
				}
			}
			j++		
		}
		i++
		
	}
	return false

}

func ReadInput(file string) (string) {
	data, _ := os.ReadFile(file)
	return string(data)
}

func WriteOutput(text string) {
	file, _ := os.Create("output.txt")
	file.WriteString(text)
}
