package controllers

import "net/http"

func GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Books"))
		if err != nil {
			return
		}
	}
}

func PostBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Books"))
		if err != nil {
			return
		}
	}
}
