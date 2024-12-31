package adapters

import (
	"go-clean/entities"
	"go-clean/usecases"

	"github.com/gofiber/fiber/v2"
)

// Input (order use case) adapter connects to Fiber's HTTP service
type HttpOrderHandler struct{
	orderUseCase usecases.OrderUseCase
}
func (s *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error{ // Performs order service
	var order entities.Order
	if err := c.BodyParser(&order); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := s.orderUseCase.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

// function to create HTTP adapter instance (which is called by user)
func NewHttpOrderHandler(orderUseCase usecases.OrderUseCase) *HttpOrderHandler{
	return &HttpOrderHandler{orderUseCase: orderUseCase}
}