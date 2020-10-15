package server

import (
	"context"
	"fmt"
	"time"
	"github.com/hashicorp/go-hclog"
	"github.com/guimunarolo/products-manage/calculator-service/models"
	protos "github.com/guimunarolo/products-manage/calculator-service/protos/calculator"
)

const (
	_blackFridayMonth, _blackFridayDay = 10, 14
	_maxDiscountPercentage = 0.1
	_initialDiscountPercentage = 0.0
)

type Calculator struct {
	log hclog.Logger
}

func NewCalculator(l hclog.Logger) *Calculator {
	return &Calculator{l}
}

func (c *Calculator) GetDiscountPercentage(u *models.User) float64 {
	pct := _initialDiscountPercentage
	today := time.Now()

	if today.Day() == _blackFridayDay && today.Month() == _blackFridayMonth {
		pct += 0.10
	}

	if today.Day() == u.DateOfBirth.Day() && today.Month() == u.DateOfBirth.Month() {
		pct += 0.05
	}

	if pct >= _maxDiscountPercentage {
		return _maxDiscountPercentage
	}

	return pct
}

func (c *Calculator) CalculateDiscountValue(p *models.Product, pct float64) int32 {
	discountValue := float64(p.PriceInCents) * pct
	
	return int32(discountValue)
}

func (c *Calculator) GetProductDiscount(ctx context.Context, rr *protos.ProductDiscountRequest) (*protos.ProductDiscountResponse, error) {
	userID := rr.GetUserId()
	productID :=  rr.GetProductId()
	c.log.Info("Handle request for GetProductDiscount", "productId", productID, "userId", userID)

	user := models.GetUserByID(userID)
	c.log.Info("Got user from DB", "User", fmt.Sprintf("%s %s", user.FirstName, user.LastName))

	product := models.GetProductByID(productID)
	c.log.Info("Got product from DB", "Product", product.Title)

	pct := c.GetDiscountPercentage(user)
	discountInCents := c.CalculateDiscountValue(product, pct)
	c.log.Info("Calculated discount", "percentage", pct, "value", discountInCents)

	return &protos.ProductDiscountResponse{Pct: pct, ValueInCents: discountInCents}, nil
}
