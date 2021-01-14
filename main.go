package main

import (
	"fmt"
	"log"
	"net/http"
	"books-list/driver"
	"books-list/models"
	"books-list/controllers"
	"database/sql"
	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB
var err error

func atTheDisco(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db = driver.Connect()
	defer db.Close()

	controller := controllers.Controller{}

	router := mux.NewRouter()
	
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")
	
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}