package main

import (
	"github.com/gin-gonic/gin"
	"github.com/working/go-clean-architecture/config"

	_entity "github.com/working/go-clean-architecture/app/entity"
	_handler "github.com/working/go-clean-architecture/app/handler"
	_repo "github.com/working/go-clean-architecture/app/repository"
)

func main() {
	r := gin.Default()
	
	config.SetupModels() // new
	db := config.GetDBConnection()
	port := config.GetPortConnection()

	repoBook := _repo.NewBookRepository(db)
	entityBook := _entity.NewBookEntity(repoBook)

	repoMeow := _repo.NewMeowRepository(db)
	entityMeow := _entity.NewMeowEntity(repoMeow)

	api := r.Group("/v1")

	_handler.NewBooksHandler(api, entityBook)
	_handler.NewMeowHandler(api, entityMeow)
	r.Run(port)
}
