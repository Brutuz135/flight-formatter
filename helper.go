package main

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filename) (Lookup [][]string, err string) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "err"
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lookup, err := reader.ReadAll()

	if err != nil {
		return "", "err"
	}

	return lookup, nil
}
