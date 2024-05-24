package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonin/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gasby", Author: "F. Scott Fitzgerald", Quantity: 4},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	// 👇 returns the status of the http request with the books beautifully indented (JSON)
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	// 👇 Using gin router
	router := gin.Default()
	// 👇 /books route defined on localhost:8080
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
