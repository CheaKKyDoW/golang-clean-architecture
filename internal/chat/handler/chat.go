package chat

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *chatHandler) HandleChat(c *fiber.Ctx) error {
	fmt.Println("Chat route hit!") // Debugging message
	return c.JSON(fiber.Map{"message": "Hello from chat!"})
}
