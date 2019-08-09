package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type (
	User struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

var db *gorm.DB

func init() {
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Debug().AutoMigrate(&User{})
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome!")
	})

	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users/{id}", oneUser).Methods("GET")
	router.HandleFunc("/users", newUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func oneUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user User
	db.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	user := User{
		Name:  name,
		Email: email,
	}
	db.Create(&user)
	fmt.Fprintf(w, "user created successfully")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := mux.Vars(r)["id"]
	db.Find(&user, id)
	email := r.PostFormValue("email")
	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "user updated successfully")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := mux.Vars(r)["id"]
	db.Find(&user, id)
	db.Delete(&user)
	fmt.Fprintf(w, "user deleted successfully")
}

func main() {
	fmt.Println("Hello World!")

	handleRequests()
}
