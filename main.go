package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux")

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

// Get All Books
func getBooks(w http.ResponseWriter, router *http.Request) {

}
// Get Single Book
func getBook(w http.ResponseWriter, router *http.Request) {
	
}
// Create a New Book
func createBook(w http.ResponseWriter, router *http.Request) {
	
}
// Update Book
func updateBook(w http.ResponseWriter, router *http.Request) {
	
}
// Delete Book
func deleteBook(w http.ResponseWriter, router *http.Request) {
	
}

func main() {
	//init router
	router := mux.NewRouter()
	//route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", getBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}