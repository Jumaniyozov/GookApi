package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jumaniyozov/semiTrashApi/models"
	"log"
	"net/http"
	"strconv"
)

func SendJson(w http.ResponseWriter, val any) {
	if err := json.NewEncoder(w).Encode(val); err != nil {
		log.Fatal("Error while serializing:", err)
	}
}

func GetBooksID(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		w.WriteHeader(400)
		msg := models.Message{Message: "Do not use paraemeter ID as uncasted to int type"}
		SendJson(w, msg)
	}

	return id
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	id := GetBooksID(w, r)

	book, ok := models.FindBookById(id)

	if !ok {
		w.WriteHeader(404)
		msg := models.Message{Message: "Book with that id doesn't exist."}
		SendJson(w, msg)
		return
	}

	log.Println("Getting book with id:", id)

	w.WriteHeader(200)
	SendJson(w, book)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		msg := models.Message{Message: "Invalid json format"}
		SendJson(w, msg)
		return
	}

	newBookID := len(models.DB) + 1
	book.ID = newBookID
	models.DB = append(models.DB, book)

	log.Println("Creating book with id:", book.ID)
	w.WriteHeader(201)
	SendJson(w, book)
}
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id := GetBooksID(w, r)

	_, ok := models.FindBookById(id)
	if !ok {
		w.WriteHeader(404)
		msg := models.Message{Message: "Book with that id doesn't exist."}
		SendJson(w, msg)
		return
	}

	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		log.Println("Error while decoding:", err)
		msg := models.Message{Message: "Invalid json format"}
		SendJson(w, msg)
		return
	}
	newBook.ID = id

	for idx, b := range models.DB {
		if b.ID == id {
			models.DB[idx] = newBook
		}
	}

	log.Println("Book with id:", id, "was updated")
	SendJson(w, newBook)
}
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id := GetBooksID(w, r)

	_, ok := models.FindBookById(id)
	if !ok {
		w.WriteHeader(404)
		msg := models.Message{Message: "Book with that id doesn't exist."}
		SendJson(w, msg)
		return
	}

	for idx, book := range models.DB {
		if book.ID == id {
			models.DB = append(models.DB[:idx], models.DB[idx+1:]...)
		}
	}

	msg := models.Message{Message: fmt.Sprintf("Book with id: %d was successfully deleted", id)}
	SendJson(w, msg)
}
