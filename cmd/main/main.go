package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/skozlovtsev/Go-MySQL-CRUD-Bookstore-Management-API/pkg/routes"
)

func main(){
	r := mux.NewRouter()  //Добавляем новый маршрутизатор
	routes.RegisterBookStoreRoutes(r)  //Передаем созданый маршрутизатор r в функцию RegisterBookStoreRoutes, для добавление обработчиков
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}