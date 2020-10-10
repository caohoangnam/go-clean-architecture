package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tinrab/retry"
	"github.com/working/go-clean-architecture/domain"
	"github.com/working/go-clean-architecture/events"
	"github.com/working/go-clean-architecture/search"
	service "github.com/working/go-clean-architecture/services"
)

var db *gorm.DB
var port string

var (
	esURL   = flag.String("url", "http://elasticsearch:9200", "Elasticsearch connection string")
	natsURL = fmt.Sprintf("http://nats:4222")
)

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
	if !db.HasTable(&domain.Meow{}) {
		db.CreateTable(&domain.Meow{})
	}
	if err != nil {
		fmt.Println("Failed to connect to database!", err)
		return
	}
	fmt.Println("Connect successfully to database!")

	db.AutoMigrate(&domain.Book{})
	db.AutoMigrate(&domain.Meow{})

	// Initialize value
	m := domain.Book{Author: "caonam", Title: "hoangnam"}

	db.Create(&m)

	SetUpDBConnection(db)
	SetPortConnection(":8080")
	// defer db.Close()

	// Initialize Elastics
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		elastic, err := search.NewElastic(esURL)
		if err != nil {
			fmt.Println("Failed to connect Elastics")
			return nil
		}
		if elastic == nil {
			fmt.Println("Failed to connect Elastics with pointer")
			return nil
		}
		fmt.Println("Connect successfully to Elastics!")
		search.SetRepository(elastic)
		return nil
	})
	// defer elastic.Close()

	//	// Initialize Nats

	//	hub := service.newHub()
	//	retry.ForeverSleep(2*time.Second, func(_ int) error {
	//		nats, err := events.NewNats(natsURL)
	//		if err != nil {
	//			fmt.Println("Failed to connect to nats", err)
	//			return nil
	//		}
	//		if nats == nil {
	//			fmt.Println("Failed to connect Nats with pointer")
	//			return nil
	//		}
	//		err = nats.OnMeowCreated(func(m events.MeowCreatedMessage) {
	//			log.Printf("Meow received: %v\n", m)
	//			hub.broadcast(newMeowCreatedMessage(m.Id, m.Body, m.CreatedAt), nil)
	//		})
	//		fmt.Println("Connect successfully to Nats!")
	//		events.SetEventStore(nats)
	//		return nil
	//	})
	//	//	// defer events.Close()
	//

	hub := service.NewHub()
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		nats, err := events.NewNats(natsURL)
		if err != nil {
			fmt.Println("Failed to connect to nats", err)
			return nil
		}
		if nats == nil {
			fmt.Println("Failed to connect Nats with pointer")
			return nil
		}

		// Push messages to clients
		err = nats.OnMeowCreated(onMeowCreated)
		if err != nil {
			log.Println(err)
			return err
		}
		err = nats.OnMeowCreated(func(m events.MeowCreatedMessage) {
			log.Printf("Meow received: %v\n", m)
			hub.Broadcast(service.NewMeowCreatedMessage(m.Id, m.Body, m.CreatedAt), nil)
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}

		events.SetEventStore(nats)
		return nil
	})
	//	defer events.Close()

	// Run WebSocket server
	go hub.Run()
	http.HandleFunc("/pusher", hub.HandleWebSocket)
	if err != nil {
		fmt.Println(err)
		return
	}
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

func onMeowCreated(m events.MeowCreatedMessage) {
	meow := domain.Meow{
		Id:        m.Id,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
	}
	if err := search.Create(context.Background(), meow); err != nil {
		log.Println(err)
	}
}
