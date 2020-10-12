package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/guimunarolo/products-manage/calculator-service/protos/calculator"
)


type Calculator struct {
	log hclog.Logger
}

func NewCalculator(l hclog.Logger) *Calculator {
	return &Calculator{l}
}

func (c *Calculator) GetProductDiscount(ctx context.Context, rr *protos.ProductDiscountRequest) (*protos.ProductDiscountResponse, error) {
	c.log.Info("Handle request for GetProductDiscount", "ProductId", rr.GetProductId(), "UserId", rr.GetUserId())
	
	return &protos.ProductDiscountResponse{Pct: 0.5, ValueInCents: 100}, nil
}
