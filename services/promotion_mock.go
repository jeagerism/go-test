package services

import "github.com/stretchr/testify/mock"

type promotionServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *promotionServiceMock {
	return &promotionServiceMock{}
}

// Correct method name: CalculateDiscount
func (m *promotionServiceMock) CalculateDiscount(amount int) (int, error) {
	args := m.Called(amount)
	return args.Int(0), args.Error(1)
}
