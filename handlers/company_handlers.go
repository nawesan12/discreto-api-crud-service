package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"discreto-api-crud-service/models"
)

func SetupCompanyRoutes(r *gin.Engine, DB *gorm.DB) {
	r.GET("/companies/:id", func(c *gin.Context) {
		GetCompany(c, DB)
	})
	// Add other routes here
}

func GetCompany(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	var company models.Company
	if err := DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}
