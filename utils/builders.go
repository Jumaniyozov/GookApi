package utils

import (
	"github.com/gorilla/mux"
	"github.com/jumaniyozov/semiTrashApi/handlers"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("DELETE")
}

func BuildManyBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
