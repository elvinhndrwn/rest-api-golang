package main

import (
	"log"

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
}
