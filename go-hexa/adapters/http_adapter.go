package adapters

import (
	"go-hexa/core"

	"github.com/gofiber/fiber/v2"
)

// Primary adapter (input)
type HttpOrderHandler struct{
	service core.OrderService
}
func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error{
	var order core.Order
	if err := c.BodyParser(&order); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler{
	return &HttpOrderHandler{service: service}
}