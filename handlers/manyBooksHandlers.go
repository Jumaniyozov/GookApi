package handlers

import (
	"github.com/jumaniyozov/semiTrashApi/models"
	"log"
	"net/http"
)

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Get infos about all books in db")

	w.WriteHeader(200)
	SendJson(w, models.DB)
}
