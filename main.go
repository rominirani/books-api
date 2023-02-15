package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rominirani/crud-mockapi/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	/*router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})*/

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("Books API is running!")
	http.ListenAndServe(":8080", router)
}
