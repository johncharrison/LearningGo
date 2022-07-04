package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, b := range books {
		if b.ID == params["id"] {
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			var book Book = item
			books = append(books[:index], books[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init mux router
	router := mux.NewRouter()

	// MOCK DATA -- TODO implement db
	books = append(books, Book{ID: "1", Isbn: "43235231", Title: "Book 1", Author: &Author{Firstname: "Frank", Lastname: "Feckers"}})
	books = append(books, Book{ID: "2", Isbn: "73847823", Title: "Book 2", Author: &Author{Firstname: "Carol", Lastname: "Askin"}})
	books = append(books, Book{ID: "3", Isbn: "91725371", Title: "Book 3", Author: &Author{Firstname: "Rick", Lastname: "Roll"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
