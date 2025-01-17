package handlers

import (
	"gotest/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionhandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionhandler{promoService: promoService}
}

func (h promotionhandler) CalculateDiscount(c *fiber.Ctx) error {
	// http://localhost:8000/calculate?amount=100
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	discount, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendString(strconv.Itoa(discount))
}
