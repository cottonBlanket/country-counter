package main

import (
	"github.com/cottonBlanket/country-counter/internal/db"
	"github.com/cottonBlanket/country-counter/internal/handler"
	"github.com/cottonBlanket/country-counter/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Инициализация БД
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	db.Init(dsn)

	userService := service.NewUserService(db.DB)

	h := handler.NewHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users", h.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/find", h.FindUsers).Methods("GET")

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
