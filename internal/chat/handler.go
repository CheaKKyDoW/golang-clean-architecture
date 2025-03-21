package chat

import (
	"github.com/gofiber/fiber/v2"
)

type ChatHandler interface {
	RegisterRoutes(r fiber.Router)
}

type chatHandler struct {
	chatUsecase ChatUsecase
}

func NewChatHandler(usecase ChatUsecase) ChatHandler {
	return &chatHandler{}
}

func (h *chatHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/chat", nil)
}
