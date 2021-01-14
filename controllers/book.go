package controllers

import (
	"books-list/models"
	"books-list/repository/book"
	"books-list/utils"
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Controller struct {}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		var error models.Error
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)
		
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		var error models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}

		book, err := bookRepo.GetBook(db, book, params["id"])	
		
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Book Not Found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		var bookID int
		var error models.Error
	
		json.NewDecoder(r.Body).Decode(&book)
	
		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
		}
	
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
		}
	
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		var error models.Error
	
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.DeleteBook(db, params["id"])

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
		}
	
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}