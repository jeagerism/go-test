package services

import "gotest/repositories"

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionServices struct {
	promotionRepo repositories.PromotionRepository
}

func NewPromotionService(promotionRepo repositories.PromotionRepository) PromotionService {
	return &promotionServices{promotionRepo: promotionRepo}
}

func (p *promotionServices) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := p.promotionRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchasesMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}
