package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // here
)

type Contact struct {
	gorm.Model

	Name    string
	Email   string `gorm:"typevarchar(100);unique_index"`
	Address string
}

var (
	contact = &Contact{
		Name: "Alex", Email: "aliaaa@gmail.com", Address: "New York",
	}
)

var db *gorm.DB
var err error

func main() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// opening connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to Database")
	}

	defer db.Close()

	// make migration to database

	db.AutoMigrate(&Contact{})

	db.Create(contact)

	// API routes

	router := mux.NewRouter()

	router.HandleFunc("/contacts", getContacts).Methods("GET")
	router.HandleFunc("/contact/{id}", getContact).Methods("GET")
	router.HandleFunc("/create/contact", createContact).Methods("POST")
	router.HandleFunc("/delete/contact/{id}", deleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contacts []Contact
	db.Find(&contacts)
	json.NewEncoder(w).Encode(&contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var contact Contact

	db.First(&contact, params["id"])

	json.NewEncoder(w).Encode(contact)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	json.NewDecoder(r.Body).Decode(&contact)

	createdContact := db.Create(&contact)
	err = createdContact.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&contact)
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var contact Contact

	db.First(&contact, params["id"])
	db.Delete(&contact)

	json.NewEncoder(w).Encode(&contact)
}
