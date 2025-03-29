package chat

import (
	handler "golang-clean-architecture/internal/chat/handler"
	repository "golang-clean-architecture/internal/chat/repository"
	usecase "golang-clean-architecture/internal/chat/usecase"

	"golang-clean-architecture/internal/infrastructure/client"
	"golang-clean-architecture/internal/infrastructure/websocket"
	"golang-clean-architecture/pkg/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitModule(cfg *config.Cfg, router fiber.Router, httpClient *client.HTTPClient, dbConn *gorm.DB, wsManager *websocket.WebSocketManager) {
	// init repository
	repository := repository.NewChatRepository(dbConn)

	// init usecase
	usecase := usecase.NewChatUsecase(repository)

	// init handler
	handler := handler.NewChatHandler(usecase, wsManager)

	handler.RegisterRoutes(router.Group("/api/v1/chat"))

}
