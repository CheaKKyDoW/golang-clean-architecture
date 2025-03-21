package chat

import (
	"golang-clean-architecture/internal/infrastructure/client"

	"golang-clean-architecture/pkg/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitModule(cfg *config.Cfg, router fiber.Router, httpClient *client.HTTPClient, dbConn *gorm.DB) {
	// init repository
	repository := NewChatRepository(dbConn)

	// init usecase
	usecase := NewChatUsecase(repository)

	// init handler
	handler := NewChatHandler(usecase)

	handler.RegisterRoutes(router.Group("/api/v1/chat"))
}
