package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type (
	User struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome!")
	})

	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users/{id}", oneUser).Methods("GET")
	router.HandleFunc("/users/{id}", newUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "all users")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func oneUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "one user")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "new user")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "update user")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "delete user")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Println("Hello World!")

	handleRequests()
}
