package main

import (
	"github.com/gin-gonic/gin"
	"github.com/working/go-clean-architecture/config"

	_bookEntity "github.com/working/go-clean-architecture/app/entity"
	_bookHandler "github.com/working/go-clean-architecture/app/handler"
	_bookRepo "github.com/working/go-clean-architecture/app/repository"
)

func main() {
	r := gin.Default()
	config.SetupModels() // new
	db := config.GetDBConnection()
	port := config.GetPortConnection()

	repo := _bookRepo.NewBookRepository(db)
	entity := _bookEntity.NewBookEntity(repo)
	api := r.Group("/v1")

	_bookHandler.NewBooksHandler(api, entity)

	r.Run(port)
}
