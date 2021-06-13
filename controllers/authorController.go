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

var myAuthorService = services.NewMyAuthorService()

func GetAuthors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")

	authors := myAuthorService.GetAuthors()

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Author successfully", authors))
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse("Please provide a valid Author id"))
		return
	}
	author, err := myAuthorService.GetAuthor(int64(id))

	if err != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Author successfully", author))
}

func PostAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)

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

	author, err1 := myAuthorService.CreateAuthor(author)
	if err1 != nil {
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err1.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Author Created successfully", author))
}
