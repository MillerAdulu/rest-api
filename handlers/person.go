package handlers

import (
	"../person"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var people []person.Person

func init() {
	people = append(people, person.Person{ID: "1", Firstname: "Miller", Lastname: "Doe", Address: &person.Address{City: "Nairobi"}})
	people = append(people, person.Person{ID: "2", Firstname: "M", Lastname: "Mzazi", Address: &person.Address{City: "Kisumu"}})
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&person.Person{})
}

func PersonsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person person.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1])
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
