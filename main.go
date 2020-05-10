package main

import (
	"github.com/gin-gonic/gin"
	"github.com/working/project/config"

	_bookHttpDelivery "github.com/working/project/book/delivery/http"
	_bookMiddleware "github.com/working/project/book/delivery/http/middleware"
	_bookRepo "github.com/working/project/book/repository/psql"
	_bookUcase "github.com/working/project/book/usecase"
)

func main() {
	r := gin.Default()
	config.SetupModels() // new
	db := config.GetDBConnection()
	port := config.GetPortConnection()

	r.Use(_bookMiddleware.Cors())

	repo := _bookRepo.NewPsqlBookRepository(db)
	us := _bookUcase.NewBookUsecase(repo)
	api := r.Group("/v1")

	_bookHttpDelivery.NewBooksHandler(api, us)

	r.Run(port)
}
