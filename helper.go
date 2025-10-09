package main

import (
	"encoding/csv"
	"os"
	"regexp"
	"time"
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

func ChangeAirportName(input string, lookupIata map[string]string, lookupIcao map[string]string) (string) {


	output := ""
	for i := 0; len(input) > i; i++ {

		found := false
		value := ""
		key := ""
		jump := 0
		if input[i] == '#' && input[i+1] == '#' && i+6 <= len(input) {
			key = input[i+2 : i+6]
			value, found = lookupIcao[key]
			jump = 6
		} else if input[i] == '#' {
			key = input[i+1 : i+4]
			value, found = lookupIata[key]
			jump = 4
		}
		if found {
			output = output + value
			i += jump
		} else {
			output = output + string(input[i])
		}

	}
	return output
}


	
	
	
func GreateMaps(airportLookup [][]string) (map[string]string, map[string]string) {
	
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

func ChangeTime(input string) (string) {
 	dre := regexp.MustCompile(`D\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[+-][0-9]{2}:[0-9]{2})\)`)
	reT12 := regexp.MustCompile(`T12\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[+-][0-9]{2}:[0-9]{2})\)`)
	reT24 := regexp.MustCompile(`T24\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[+-][0-9]{2}:[0-9]{2})\)`)

	output := dre.ReplaceAllStringFunc(input, DateChange)
	output = reT12.ReplaceAllStringFunc(output, T12Change)
	output = reT24.ReplaceAllStringFunc(output, T24Change)
	return output
	

}

func T24Change(input string) (string) {
	input = input[4:len(input)-1]
	layout := "2006-01-02T15:04Z07:00"

	t, _ := time.Parse(layout, input)
	newDate := t.Format("15:04 (-07:00)")
	return newDate
}



func T12Change(input string) (string) {
	input = input[4:len(input)-1]
	layout := "2006-01-02T15:04Z07:00"

	t, _ := time.Parse(layout, input)
	newDate := t.Format("03:04PM (-07:00)")
	return newDate
}

func DateChange(input string) (string) {
	input = input[2:len(input)-1] 
	layout := "2006-01-02T15:04Z07:00"

	t, _ := time.Parse(layout, input)
	newDate := t.Format("02 Jan 2006")
	return newDate
}

func TrimSpaces(input string) (string) {
	re := regexp.MustCompile(`[\v\f\r]`)
	input = re.ReplaceAllString(input, "\n")

	ren := regexp.MustCompile(`\n{3,}`)
	input = ren.ReplaceAllString(input, "\n\n")
	return input
}
	

