package main

import (
	"./handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", handlers.PersonsHandler).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.PersonHandler).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePersonsHandler).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePersonHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
