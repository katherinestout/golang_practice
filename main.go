package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv")

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

// Book var as a slice Book struct (array kinda)
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
// Get Single Book
func getBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) //get params
	//loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
// Create a New Book
func createBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_= json.NewDecoder(router.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //generate mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	
}
// Update Book
func updateBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_= json.NewDecoder(router.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}	
	}
	json.NewEncoder(w).Encode(books)
}
// Delete Book
func deleteBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}	
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init router
	router := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "123456", Title: "Book One", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "123457", Title: "Book Two", Author: &Author {Firstname: "Jenny", Lastname: "Smith"}})
	//route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}