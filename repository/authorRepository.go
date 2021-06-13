package repository

import (
	"errors"
	"restProject/db"
	"restProject/models"
	"strconv"
)

func GetAuthors() []models.Author {
	authors := make([]models.Author, 0)
	for _, author := range db.Authors {
		authors = append(authors, author)
	}

	return authors
}

func CreateAuthor(a models.Author) error {
	authors := db.Authors
	if _, ok := authors[a.GetID()]; ok {
		return errors.New("Author with ID:" + strconv.Itoa(int(a.GetID())) + " already exists in db")
	}

	authors[a.GetID()] = a
	return nil
}

func GetAuthor(id int64) (models.Author, error) {
	authors := db.Authors
	_, ok := authors[id]
	if !ok {
		return models.Author{}, errors.New("Author with ID:" + strconv.Itoa(int(id)) + " doesn't exist in db")
	}

	return authors[id], nil
}
