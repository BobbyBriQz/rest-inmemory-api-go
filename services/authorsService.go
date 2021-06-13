package services

import (
	"restProject/models"
	"restProject/repository"
	"sort"
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
	authors := repository.GetAuthors()
	sort.Slice(authors, func(i, j int) bool {
		return authors[i].GetID() < authors[j].GetID()
	})

	return authors
}

func (m MyAuthorService) CreateAuthor(b models.Author) (models.Author, error) {
	id := len(repository.GetAuthors())
	b.SetID(int64(id) + 1)

	return repository.CreateAuthor(b)
}

func (m MyAuthorService) GetAuthor(id int64) (models.Author, error) {
	return repository.GetAuthor(id)
}
