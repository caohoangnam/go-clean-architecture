package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/working/go-clean-architecture/domain"
)

type BookeHandler struct {
	BookEntity domain.BookEntity
}

func NewBooksHandler(r *gin.RouterGroup, us domain.BookEntity) {
	handler := &BookeHandler{
		BookEntity: us,
	}
	r.GET("/books", handler.FindBooks)
	r.GET("/books/:id", handler.FindBook)
}

func (a *BookeHandler) FindBooks(c *gin.Context) {
	books, _ := a.BookEntity.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": books})
}

func (a *BookeHandler) FindBook(c *gin.Context) {
	books, _ := a.BookEntity.GetByID(c.Request.Context(), c.Param("id"))
	c.JSON(200, gin.H{"data": books})
}
