package http

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
	// r.POST("/books", handler.CreateBook)       // create
	r.GET("/books/:id", handler.FindBook) // find by id
	// r.PATCH("/books/:id", handler.UpdateBook)  // update by id
	// r.DELETE("/books/:id", handler.DeleteBook) // delete by id
}

func (a *BookeHandler) FindBooks(c *gin.Context) {
	books, _ := a.BookUsecase.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": books})
}

func (a *BookeHandler) FindBook(c *gin.Context) {
	books, _ := a.BookUsecase.GetByID(c.Request.Context(), c.Param("id"))
	c.JSON(200, gin.H{"data": books})
}

// // POST /books
// // Create new books
// func CreateBook(c *gin.Context) {
// 	// Validate input
// 	var input models.CreateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// Create Book
// 	book := models.Book{Title: input.Author, Author: input.Author}
// 	db.Create(&book)
// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // GET /books/:id
// // Find a book
// func FindBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // PATCH /books/:id
// // Update a book
// func UpdateBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	// Validate input
// 	var input models.UpdateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	db.Model(&book).Updates(input)
// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // DELETE /books/:id
// // Delete a book
// func DeleteBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	db.Delete(&book)
// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
