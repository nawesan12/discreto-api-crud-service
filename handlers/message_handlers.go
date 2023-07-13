package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"discreto-api-crud-service/models"
)

func SetupMessageRoutes(r *gin.Engine, DB *gorm.DB) {
	r.GET("/messages/:id", func(c *gin.Context) {
		GetMessage(c, DB)
	})
	// Add other routes here
}

func GetMessage(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	var message models.Message
	if err := DB.First(&message, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}
	c.JSON(http.StatusOK, message)
}
