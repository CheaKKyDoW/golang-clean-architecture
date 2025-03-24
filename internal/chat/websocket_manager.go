package chat

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// WebSocketManager manages all WebSocket connections
type WebSocketManager struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

// NewWebSocketManager initializes WebSocket manager
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[*websocket.Conn]bool),
	}
}

// HandleConnection processes WebSocket connections
func (wm *WebSocketManager) HandleConnection(c *websocket.Conn) {
	defer func() {
		wm.mu.Lock()
		delete(wm.clients, c)
		wm.mu.Unlock()
		c.Close()
	}()

	wm.mu.Lock()
	wm.clients[c] = true
	wm.mu.Unlock()
	log.Println("New WebSocket client connected")

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket disconnected:", err)
			break
		}

		fmt.Printf("Received message: %s\n", msg)

		// Broadcast message to all clients
		wm.Broadcast(msg)
	}
}

// Broadcast sends messages to all connected clients
func (wm *WebSocketManager) Broadcast(msg []byte) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	for client := range wm.clients {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error broadcasting:", err)
		}
	}
}
