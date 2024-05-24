package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
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

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	// get the id from the query param
	id, ok := c.GetQuery("id")

	// checking if id isn't missing
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	book, err := getBookById(id)

	// checking if a book with this id exists
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Didn't find a book with this id."})
		return
	}

	// checking if the book's available
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book unavailable..."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		// return a message when not found
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Didn't find a book with this id."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book
	//  bind data of the request to the newBook created above
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// append newBook to books
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	// Using gin router
	router := gin.Default()
	// /books route defined on localhost:8080
	router.GET("/books", getBooks)
	router.GET("/book/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
