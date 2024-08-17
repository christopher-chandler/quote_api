package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	data "quotientary/parser"
)

func GinGetAllQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.CsvQuotes("resources/medium_quotes.csv"))
}

func GinGetQuoteByText(c *gin.Context) {
	text := c.Param("text")
	quote, err := GetQuoteByText(text)
	fmt.Println(text, quote)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Text not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, quote)
}
func GinGetQuoteByAuthor(c *gin.Context) {
	author := c.Param("author")
	quote, err := GetQuoteByAuthor(author)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Author not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, quote)
}

func GinQuoteAuthorCategory(c *gin.Context) {
	// Retrieve query parameters
	quote := c.Query("quote")
	author := c.Query("author")
	category := c.Query("category")

	_, err := GetQuoteByAuthor(author)

	// If you want to use a default value if the query parameter is not present
	default_query := c.DefaultQuery("default", "default_value")

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Author not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"quote":    quote,
		"author":   author,
		"category": category,
		"default":  default_query,
	})

}
