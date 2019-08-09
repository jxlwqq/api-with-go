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
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Hello World!")

	handleRequests()
}
