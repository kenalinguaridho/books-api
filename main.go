package main

import (
	"github.com/kenalinguaridho/books-api/config"
	"github.com/kenalinguaridho/books-api/controller/authcontroller"
	"github.com/kenalinguaridho/books-api/controller/authorcontroller"
	"github.com/kenalinguaridho/books-api/controller/bookcontroller"
	"github.com/kenalinguaridho/books-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	config.DBConnect()

	// Route handler login
	router.HandleFunc("/auth/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/auth/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/auth/logout", authcontroller.Logout).Methods("GET")
	
	api := router.PathPrefix("/api").Subrouter()

	// route handler author
	api.HandleFunc("/authors", authorcontroller.Index).Methods("GET")
	api.HandleFunc("/authors", authorcontroller.Create).Methods("POST")
	api.HandleFunc("/authors/{id}/detail", authorcontroller.Detail).Methods("GET")
	api.HandleFunc("/authors/{id}/update", authorcontroller.Update).Methods("PUT")
	api.HandleFunc("/authors/{id}/delete", authorcontroller.Delete).Methods("DELETE")
		
	// route handler books
	api.HandleFunc("/books", bookcontroller.Index).Methods("GET")
	api.HandleFunc("/books", bookcontroller.Create).Methods("POST")
	api.HandleFunc("/books/{id}/detail", bookcontroller.Detail).Methods("GET")
	api.HandleFunc("/books/{id}/update", bookcontroller.Update).Methods("PUT")
	api.HandleFunc("/books/{id}/delete", bookcontroller.Delete).Methods("DELETE")
	
	api.Use(middleware.JWTMiddleware)

	log.Println("Server running at port 8080")
	http.ListenAndServe(":8080", router)
}