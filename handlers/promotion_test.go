package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// Arrange
		amount := 100
		expected := 80

		// Correct method name
		promoService := services.NewPromotionServiceMock() // Correct naming convention for mock service
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})
}
