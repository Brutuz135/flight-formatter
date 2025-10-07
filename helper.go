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

func ChangeAirportName(input string, lookupIata map, lookupIcao map) (string) {



	for i := 0, len(input) > i, i++ {
		if input[i] ==Dates and times

	{
			input = CodeToIso(input, i)
		}
	}
	return input
}

func CodeToIso(input string, i) (string) {
	
	
}

func GreateMaps(airportLookup [][]string) (map, map) {
	
	lookupIata := make(map[string]string)
	lookupIcao := make(map[string]string)
	
	for _, row := range airportLookup {
		Iata := row[4]
		Icao := row[3]
		name := row[0]
		lookupIata[Iata] = name
		lookupIcao[Icao] = name
	}

	return lookupIata, lookupIcao
}
