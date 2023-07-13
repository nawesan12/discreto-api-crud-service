package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"discreto-api-crud-service/handlers"
)

func SetupRoutes(r *gin.Engine, DB *gorm.DB) {
	handlers.SetupUserRoutes(r, DB)
	handlers.SetupCompanyRoutes(r, DB)
	handlers.SetupChatRoomRoutes(r, DB)
	handlers.SetupPrivateChatRoutes(r, DB)
	handlers.SetupMessageRoutes(r, DB)
}
