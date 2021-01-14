package bookRepository

import (
	"books-list/models"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from books")
	if err != nil {
		return []models.Book{}, err
	}
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, p string) (models.Book, error) {
	row := db.QueryRow("select * from books where id = $1", p)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	var bookID int
	err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) returning id;", 
		book.Title, book.Author, book.Year).Scan(&bookID)
	if err != nil {
		return 0, err
	}

	return bookID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	res, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) DeleteBook(db *sql.DB, p string) (int64, error) {
	res, err := db.Exec("delete from books where id=$1", p)
	if err != nil {
		return 0, err
	}
	
	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}