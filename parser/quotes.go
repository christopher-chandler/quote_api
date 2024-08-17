package parser

import (
	"quotientary/src/body"
)

func CsvQuotes(CsvPath string) []body.Quote {
	/*
		Args:
			CsvPath (string): The path to the CSV file.

		Returns:
			[]body.Quote: A slice of body.Quote objects
			where each object represents a quote
			extracted from the CSV file.
	*/
	records := ReadCsvFile(CsvPath)
	var csvQuoteData []body.Quote

	for i := 1; i < len(records); i++ {
		var data = records[i]
		csvQuoteData = append(csvQuoteData, body.Quote{
			Text:     data[0],
			Author:   data[1],
			Category: data[2],
		})
	}

	return csvQuoteData

}
