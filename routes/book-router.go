package router

import (
	"book-management-system/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/create", controllers.CreateBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/books/update/{id}", controllers.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/books/delete/{id}", controllers.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
