package routes

import (
	"github.com/gorilla/mux"
	"github.com/skozlovtsev/Go-MySQL-CRUD-Bookstore-Management-API/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//добавление обработчиков запросов
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
