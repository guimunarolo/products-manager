package calculator

import (
	"context"
	"time"

	"github.com/hashicorp/go-hclog"
)

const (
	blackFridayMonth, blackFridayDay = 11,  25
	maxDiscountPercentage = 0.1
	initialDiscountPercentage = 0.0
)

type Calculator struct {
	logger            hclog.Logger
	userRepository    UserRespository
	productRepository ProductRepository
}

func NewCalculator(logger hclog.Logger, urep UserRespository, prep ProductRepository) *Calculator {
	return &Calculator{
		logger:            logger,
		userRepository:    urep,
		productRepository: prep,
	}
}

func (calc *Calculator) GetProductDiscount(ctx context.Context, pdr *ProductDiscountRequest) (*ProductDiscountResponse, error) {
	userID := pdr.GetUserId()
	productID :=  pdr.GetProductId()
	
	calc.logger.Info("Handle request for GetProductDiscount", "product_id", productID, "user_id", userID)

	user, err := calc.userRepository.GetUser(userID)
	if err != nil {
		calc.logger.Error("Could not find this user")
		return nil, err
	}

	product,err := calc.productRepository.GetProduct(productID)
	if err != nil {
		calc.logger.Error("Could not find this product")
		return nil, err
	}

	pct := getDiscountPercentage(user)
	discountInCents := calculateDiscountValue(product, pct)
	calc.logger.Info("Calculated discount", "percentage", pct, "value", discountInCents)

	return &ProductDiscountResponse{Pct: pct, ValueInCents: discountInCents}, nil
}

func getDiscountPercentage(u *User) float64 {
	pct := initialDiscountPercentage
	today := time.Now()

	if today.Day() == blackFridayDay && today.Month() == blackFridayMonth {
		pct += 0.10
	}

	if today.Day() == u.DateOfBirth.Day() && today.Month() == u.DateOfBirth.Month() {
		pct += 0.05
	}

	if pct >= maxDiscountPercentage {
		return maxDiscountPercentage
	}

	return pct
}

func calculateDiscountValue(p *Product, pct float64) int32 {
	discountValue := float64(p.PriceInCents) * pct
	
	return int32(discountValue)
}
