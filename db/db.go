package db

import "restProject/models"

var Books = make(map[int64]models.Book)
var Authors = make(map[int64]models.Author)
