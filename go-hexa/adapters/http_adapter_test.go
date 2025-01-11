package adapters

import (
	"bytes" // convert to bytes
	"errors"
	"go-hexa/core"
	"net/http/httptest" // create request
	"testing"

	"github.com/gofiber/fiber/v2"        // http server
	"github.com/stretchr/testify/assert" // expected results comparison
	"github.com/stretchr/testify/mock"   // mock functions
)

// MockOrderService is a mock implementation of core.OrderService
type mockOrderService struct {
  mock.Mock
}

func (m *mockOrderService) CreateOrder(order core.Order) error {
  args := m.Called(order)
  return args.Error(0)
}

// TestCreateOrderHandler tests the CreateOrder handler of HttpOrderHandler
func TestCreateOrderHandler(t *testing.T) {
  mockService := new(mockOrderService)
  handler := NewHttpOrderHandler(mockService)

  app := fiber.New()
  app.Post("/orders", handler.CreateOrder)

  // Test case: Successful order creation
  t.Run("Successful order creation", func(t *testing.T) {
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(nil) // mock CreateOrder function

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": 100}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
    mockService.AssertExpectations(t) // mock.On()/Return() was in fact called
  })

  t.Run("Fail order creation (total less than 0)", func(t *testing.T) {
    mockService.ExpectedCalls = nil // mockService.ExpectedCalls is the status if the call functions in the object(struct) were called
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("total must be positive"))

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": -200}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
    mockService.AssertExpectations(t) // mock.On()/Return() was in fact called
  })

  // Test case: Invalid request body
  t.Run("Invalid request body", func(t *testing.T) {
    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": "invalid"}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
  })

  // Test case: Order service returns an error
  t.Run("Order service error", func(t *testing.T) {
    mockService.ExpectedCalls = nil // mockService.ExpectedCalls is the status if the call functions in the object(struct) were called
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("service error"))

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": 100}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
    mockService.AssertExpectations(t) // mock.On()/Return() was in fact called
  })
}

