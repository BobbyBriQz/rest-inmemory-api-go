package repository

import (
	"errors"
	"restProject/db"
	"restProject/models"
	"strconv"
)

func GetBooks() []models.Book {
	books := make([]models.Book, 0)
	for _, book := range db.Books {
		books = append(books, book)
	}

	return books
}

func CreateBook(b models.Book) (models.Book, error) {
	books := db.Books
	if _, ok := books[b.ID]; ok {
		return models.Book{}, errors.New("Book with ID:" + strconv.Itoa(int(b.ID)) + " already exists in db")
	}

	books[b.ID] = b
	return b, nil
}

func GetBook(id int64) (models.Book, error) {
	books := db.Books
	_, ok := books[id]
	if !ok {
		return models.Book{}, errors.New("Book with ID:" + strconv.Itoa(int(id)) + " doesn't exist in db")
	}

	return books[id], nil
}
