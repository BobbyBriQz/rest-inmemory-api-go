package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"restProject/models"
	"restProject/services"
	"strconv"
)

var myBookService = services.NewMyBookService()

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	books := myBookService.GetBooks()

	_ = json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		_, _ = w.Write([]byte("Please provide the Book id"))
		return
	}
	book, err := myBookService.GetBook(int64(id))

	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err1 := json.NewEncoder(w).Encode(&book)
	if err1 != nil {
		_, _ = w.Write([]byte(err1.Error()))
		return
	}
}

func PostBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err1 := myBookService.CreateBook(book)
	if err1 != nil {
		_, _ = w.Write([]byte(err1.Error()))
		return
	}

	_, _ = w.Write([]byte("Book Added Successfully"))
}
