package services

import (
	"restProject/models"
	"restProject/repository"
)

type AuthorService interface {
	GetBooks() []models.Book
	CreateBook(b models.Book) error
	GetBook(id int64) (models.Book, error)
}

type MyAuthorService struct {
}

func NewMyAuthorService() MyAuthorService {
	return MyAuthorService{}
}

func (m MyAuthorService) GetAuthors() []models.Author {

	return repository.GetAuthors()
}

func (m MyAuthorService) CreateAuthor(b models.Author) error {
	id := len(repository.GetAuthors())
	b.SetID(int64(id) + 1)

	return repository.CreateAuthor(b)
}

func (m MyAuthorService) GetAuthor(id int64) (models.Author, error) {
	return repository.GetAuthor(id)
}
