package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name  string
	Email string
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int
}

var db *gorm.DB
var err error

func main() {
	// Loading enviroment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	// Close the databse connection when the main function closes
	defer db.Close()

	// Make migrations to the database if they haven't been made already
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
}
