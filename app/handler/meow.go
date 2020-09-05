package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/working/go-clean-architecture/domain"
	"github.com/working/go-clean-architecture/events"
)

type MeowHandle struct {
	MeowEntity domain.MeowEntity
}

func NewMeowHandler(r *gin.RouterGroup, dme domain.MeowEntity) {
	handler := &MeowHandle{
		MeowEntity: dme,
	}
	r.GET("/meows", handler.List)
	r.POST("/meows", handler.Create)
}

func (m *MeowHandle) Create(c *gin.Context) {
	var meow domain.Meow
	if err := c.ShouldBindJSON(&meow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := m.MeowEntity.Create(c.Request.Context(), meow)
	if err != nil {
		fmt.Println("Can't create Meows", err)
		return
	}

	//	utils.ResponseSuccess()

	// Publish event
	err = events.PublishMeowCreated(meow)
	if err != nil {
		fmt.Println("Can't push event create meow", err)
		return
	}

	c.JSON(200, gin.H{"data": meow})
}

func (m *MeowHandle) List(c *gin.Context) {
	skip, _ := strconv.ParseInt(c.Param("skip"), 10, 64)
	take, _ := strconv.ParseInt(c.Param("take"), 10, 64)
	fmt.Println("skip", skip, "---take", take)
	meows, _ := m.MeowEntity.List(c.Request.Context(), skip, take)
	c.JSON(200, gin.H{"data": meows})
}
