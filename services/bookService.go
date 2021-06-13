package services

import (
	"restProject/models"
	"restProject/repository"
)

type BookService interface {
	GetBooks() []models.Book
	CreateBook(b models.Book) error
	GetBook(id int64) (models.Book, error)
}

type MyBookService struct {
}

func NewMyBookService() MyBookService {
	return MyBookService{}
}

func (m MyBookService) GetBooks() []models.Book {

	return repository.GetBooks()
}

func (m MyBookService) CreateBook(b models.Book) error {
	id := len(repository.GetBooks())
	b.SetID(int64(id) + 1)

	return repository.CreateBook(b)
}

func (m MyBookService) GetBook(id int64) (models.Book, error) {
	return repository.GetBook(id)
}
