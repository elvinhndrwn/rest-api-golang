package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// For table Booking
type Product struct {
	Id    int    `json:”id”`
	Code  string `json:”code”`
	Name  string `json:”name”`
	Price int    `json:"price" sql:"decimal(16,2)"`
}

// For result in array
type Result struct {
	Code    int         `json:”code”`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	// Please define your username and password for MySQL.
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar-golang?charset=utf8&parseTime=True")
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&Product{})

	handleRequests()
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8080/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/products", createProduct).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Product")

	// Tangkap payload body dari request, _ artinya tidak perlu menangkap errornya
	payloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payloads, &product)

	// Simpan ke DB
	db.Create(&product)

	// Set Response
	res := Result{Code: 200, Data: product, Message: "Success"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}
