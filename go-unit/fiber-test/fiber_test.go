package fiber

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestUserRoute function for testing the /users route.
func TestUserRoute(t *testing.T) {
  app := setup()

  // Define test cases
  tests := []struct {
    description  string
    requestBody  User
    expectStatus int
  }{
    {
      description:  "Valid input",
      requestBody:  User{"jane.doe@example.com", "Jane Doe", 30},
      expectStatus: fiber.StatusOK,
    },
    {
      description:  "Invalid email",
      requestBody:  User{"invalid-email", "Jane Doe", 30},
      expectStatus: fiber.StatusBadRequest,
    },
    {
      description:  "Invalid fullname",
      requestBody:  User{"jane.doe@example.com", "12345", 30},
      expectStatus: fiber.StatusBadRequest,
    },
    {
      description:  "Invalid age",
      requestBody:  User{"jane.doe@example.com", "Jane Doe", -5},
      expectStatus: fiber.StatusBadRequest,
    },
  }

  // Run tests
  for _, test := range tests {
    t.Run(test.description, func(t *testing.T) {
      reqBody, _ := json.Marshal(test.requestBody) // convert go struct into JSON
      req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
      req.Header.Set("Content-Type", "application/json")
      resp, _ := app.Test(req)

      assert.Equal(t, test.expectStatus, resp.StatusCode) // compare expected result and actual result
    })
  }
}