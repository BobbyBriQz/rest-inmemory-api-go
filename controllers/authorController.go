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

type AuthorController struct {
	as services.AuthorService
}

func NewAuthorController(as services.AuthorService) *AuthorController {
	return &AuthorController{as: as}
}

func (a *AuthorController) GetAuthors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")

	authors := a.as.GetAuthors()

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Author successfully", authors))
}

func (a *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse("Please provide a valid Author id"))
		return
	}
	author, err := a.as.GetAuthor(int64(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Retrieved Author successfully", author))
}

func (a *AuthorController) PostAuthor(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	author, err = a.as.CreateAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.NewAPIFailedResponse(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(models.NewAPISuccessResponse("Author Created successfully", author))
}
