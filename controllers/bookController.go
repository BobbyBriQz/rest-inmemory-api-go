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

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")

	books := myBookService.GetBooks()

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Book successfully", books))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse("Please provide a valid Book id"))
		return
	}
	book, err := myBookService.GetBook(int64(id))

	if err != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Book successfully", book))
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}

	}(r.Body)

	if err != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	book, err1 := myBookService.CreateBook(book)
	if err1 != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err1.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Book Created successfully", book))
}
