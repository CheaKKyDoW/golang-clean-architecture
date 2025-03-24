package websocket

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WebSocketManager struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (wm *WebSocketManager) HandleWebSocket(c *websocket.Conn) {
	defer c.Close()

	wm.mu.Lock()
	wm.clients[c] = true
	wm.mu.Unlock()

	log.Println("New WebSocket connection")

	for {
		messageType, msg, err := c.ReadMessage()
		if err != nil {
			break
		}

		log.Printf("Received: %s", msg)

		wm.mu.Lock()
		for client := range wm.clients {
			if err := client.WriteMessage(messageType, msg); err != nil {
				client.Close()
				delete(wm.clients, client)
			}
		}
		wm.mu.Unlock()
	}

	wm.mu.Lock()
	delete(wm.clients, c)
	wm.mu.Unlock()

	log.Println("WebSocket connection closed")
}
