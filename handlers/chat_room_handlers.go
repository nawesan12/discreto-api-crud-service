package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"discreto-api-crud-service/models"
)

func SetupChatRoomRoutes(r *gin.Engine, DB *gorm.DB) {
	r.GET("/users/:id", func(c *gin.Context) {
		GetChatRoom(c, DB)
	})
	// Add other routes here
}

func GetChatRoom(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	var chatRoom models.ChatRoom
	if err := DB.First(&chatRoom, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat room not found"})
		return
	}
	c.JSON(http.StatusOK, chatRoom)
}
