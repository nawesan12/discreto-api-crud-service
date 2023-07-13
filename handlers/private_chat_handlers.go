package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"discreto-api-crud-service/models"
)

func SetupPrivateChatRoutes(r *gin.Engine, DB *gorm.DB) {
	r.GET("/chats/:id", func(c *gin.Context) {
		GetPrivateChat(c, DB)
	})
	// Add other routes here
}

func GetPrivateChat(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	var chat models.PrivateChat
	if err := DB.First(&chat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat not found"})
		return
	}
	c.JSON(http.StatusOK, chat)
}
