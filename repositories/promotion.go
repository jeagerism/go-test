package repositories

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

type Promotion struct {
	ID              int
	PurchasesMin    int
	DiscountPercent int
}
