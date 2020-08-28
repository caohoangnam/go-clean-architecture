package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/working/go-clean-architecture/domain"
)

var db *gorm.DB
var port string

func SetupModels() {

	connectionParams := "user=docker password=docker sslmode=disable host=db"
	db, err := gorm.Open("postgres", connectionParams)
	if err != nil {
		return
	}

	// create table if it does not exist
	if !db.HasTable(&domain.Book{}) {
		db.CreateTable(&domain.Book{})
	}

	if err != nil {
		fmt.Println("Failed to connect to database!", err)
	}
	fmt.Println("Connect successfully to database!")

	db.AutoMigrate(&domain.Book{})

	// Initialize value
	m := domain.Book{Author: "caonam", Title: "hoangnam"}

	db.Create(&m)

	SetUpDBConnection(db)
	SetPortConnection(":8080")
}

func SetUpDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}

func SetPortConnection(Port string) {
	port = Port
}

func GetPortConnection() string {
	return port
}
