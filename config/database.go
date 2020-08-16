package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var db *gorm.DB
var port string

func init() {
	viper.SetConfigFile(`.env`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func SetupModels() {

	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetInt(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASSWORD`)
	dbName := viper.GetString(`DB_NAME`)

	portServer := viper.GetString(`PORT_SERVER`)

	prosgret_conname := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := gorm.Open("postgres", prosgret_conname)

	if err != nil {
		fmt.Println("Failed to connect to database!", err)
	}
	fmt.Println("Connect successfully to database!")

	// db.AutoMigrate(&domain.Book{})

	// Initialize value
	// m := domain.Book{Author: "author2", Title: "title2"}

	// db.Create(&m)

	SetUpDBConnection(db)
	SetPortConnection(portServer)
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
