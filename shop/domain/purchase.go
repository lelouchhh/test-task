package domain

import "context"

type Purchase struct {
	UserId string `json:"user_id"  validate:"required"`
	AlgoId string `json:"algo_id"  validate:"required"`
}
type PurchaseUsecase interface {
	BuyPurchase(ctx context.Context, p Purchase) (Algorithm, error)
}

type PurchaseRepository interface {
	BuyPurchase(ctx context.Context, p Purchase) (Algorithm, error)
}
