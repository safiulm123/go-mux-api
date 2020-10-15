package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"	
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"

)

// Book Struct (Model)
type Book struct{
	ID	string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}




//Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`

}


// Init Books var as a slice Book struct
var books []Book

// GET All Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)

}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r) // get Params
	//Loop through books and find with id
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
	
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	var book Book
	fmt.Println(book)
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)

	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){

	
}

func deleteBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index, item:= range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
	
}

func main (){
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement Databse
	books = append(books, Book{ID: "1", Isbn:"121212", Title:"Book One", Author: &Author{Firstname:"John", Lastname:"Joe"}})
	books = append(books, Book{ID: "2", Isbn:"29012", Title:"Book 2", Author: &Author{Firstname:"Steve", Lastname:"Smith"}})

	// Route Handlers / Endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))
}