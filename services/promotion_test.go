package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchasesMin:    c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promotionService := services.NewPromotionService(promoRepo)

			//Act
			discount, _ := promotionService.CalculateDiscount(c.amount)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchasesMin:    100,
			DiscountPercent: 20,
		}, nil)

		promotionService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promotionService.CalculateDiscount(0)

		//Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))

		promotionService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promotionService.CalculateDiscount(100)

		//Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
