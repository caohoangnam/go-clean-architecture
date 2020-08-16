package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/working/go-clean-architecture/domain"
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
	meow := m.MeowEntity.Create(c.Request.Context(), domain.Meow{})
	c.JSON(200, gin.H{"data": meow})
}

func (m *MeowHandle) List(c *gin.Context) {
	skip, _ := strconv.ParseInt(c.Param("skip"), 10, 64)
	take, _ := strconv.ParseInt(c.Param("take"), 10, 64)
	meows, _ := m.MeowEntity.List(c.Request.Context(), skip, take)
	c.JSON(200, gin.H{"data": meows})
}
