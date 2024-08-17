package api

import (
	"errors"
	"fmt"
	data "quotientary/parser"
	"quotientary/src/body"
	"regexp"
	"strings"
)

var CsvQuotes []body.Quote = data.CsvQuotes("corpus/medium_quotes.csv")

func GetQuoteByCategory() {

}

func GetQuoteByText(quoteText string) (*[]body.Quote, error) {
	var results []body.Quote

	var incomingQuote string = strings.Replace(quoteText, "_", " ", 10)

	for i := range CsvQuotes {
		var structQuote string = CsvQuotes[i].Text

		match, _ := regexp.MatchString(incomingQuote, structQuote)
		fmt.Println(match)
		if match {
			results = append(results, CsvQuotes[i])
		}
	}

	if len(results) > 0 {
		return &results, nil
	}

	return nil, errors.New("No quote found.")

}

func GetQuoteByAuthor(author string) (*[]body.Quote, error) {

	var results []body.Quote

	for i := range CsvQuotes {
		if CsvQuotes[i].Author == strings.Replace(author, "_", " ", 10) {
			results = append(results, CsvQuotes[i])
		}
	}

	if len(results) > 0 {
		return &results, nil
	}

	return nil, errors.New("author not found")
}
