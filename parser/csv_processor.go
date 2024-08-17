package parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file  "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to parse file as CSV for '%v'\n", filePath)
		log.Fatal(err)
	}

	return records
}

func SaveCsvFile(filePath string, records [][]string, saveAmount int) {

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Unable to create output file  "+filePath, err)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	var c int = 0

	for _, record := range records {
		_ = csvWriter.Write(record)
		c++
		if c == saveAmount {
			break
		}

	}
	csvWriter.Flush()

}
