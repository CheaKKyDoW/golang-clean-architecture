package chat

import (
	usecase "golang-clean-architecture/internal/chat/usecase"
	"golang-clean-architecture/internal/infrastructure/websocket"
	"log"

	"github.com/gofiber/contrib/websocket"

	"github.com/gofiber/fiber/v2"
)

type ChatHandler interface {
	RegisterRoutes(r fiber.Router)
}

type chatHandler struct {
	chatUsecase usecase.ChatUsecase
	wsManager   *websocket.WebSocketManager
}

func NewChatHandler(usecase usecase.ChatUsecase, wsManager *websocket.WebSocketManager) ChatHandler {
	return &chatHandler{
		chatUsecase: usecase,
		wsManager:   wsManager,
	}
}

func (h *chatHandler) RegisterRoutes(r fiber.Router) {

	r.Get("/chat", h.HandleChat)

	r.Use("/ws", h.wsConn)

	r.Get("/ws/:id", websocket.New(h.handleWebSocket))

}

func (h *chatHandler) wsConn(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
func (h *chatHandler) handleWebSocket(c *websocket.Conn) {
	defer c.Close()

	h.wsManager.AddClient(c)
	log.Println("New WebSocket connection:", c.Params("id"))

	for {
		messageType, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		log.Printf("Received message: %s", msg)

		h.wsManager.Broadcast(messageType, msg)
	}

	h.wsManager.RemoveClient(c)
}
