package main

import (
	"github.com/gin-gonic/gin"
	"quotientary/src/api"
)

func run_api() {

	r := gin.Default()

	r.GET("/author/:author", api.GinGetQuoteByAuthor)
	r.GET("/text/:text", api.GinGetQuoteByText)
	r.GET("/all-quotes", api.GinGetAllQuotes)
	//r.GET("/search", api.GinQuoteAuthorCategory)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func main() {

	run_api()

}
