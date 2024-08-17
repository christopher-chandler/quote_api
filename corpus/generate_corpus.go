package main

import (
	"fmt"
	"quotientary/parser"
)

func main() {

	var csv_file = parser.ReadCsvFile("corpus/quotes.csv")

	parser.SaveCsvFile("corpus/medium_quotes.csv", csv_file,
		300_000)

	fmt.Println("Corpus Generated")
}
