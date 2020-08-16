package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/working/project/domain"
)

type BookeHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBooksHandler(r *gin.RouterGroup, us domain.BookUsecase) {
	handler := &BookeHandler{
		BookUsecase: us,
	}
	r.GET("/books", handler.FindBooks)
	r.GET("/books/:id", handler.FindBook)
}

func (a *BookeHandler) FindBooks(c *gin.Context) {
	books, _ := a.BookUsecase.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": books})
}

func (a *BookeHandler) FindBook(c *gin.Context) {
	books, _ := a.BookUsecase.GetByID(c.Request.Context(), c.Param("id"))
	c.JSON(200, gin.H{"data": books})
}
