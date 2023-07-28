package usecase

import (
	"context"
	"fmt"
	"shop/domain"
	"time"
)

type PurchaseUsecase struct {
	PurchaseRepo domain.PurchaseRepository
	UserRepo     domain.UserRepository

	contextTimeout time.Duration
}

func (p PurchaseUsecase) BuyPurchase(ctx context.Context, purchase domain.Purchase) (domain.Algorithm, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	algo, err := p.PurchaseRepo.BuyPurchase(ctx, purchase)
	if err != nil {
		return domain.Algorithm{}, err
	}
	debt := domain.Debt{
		Id:   purchase.UserId,
		Debt: algo.Price,
	}
	fmt.Println(debt, algo)
	err = p.UserRepo.IncreaseDebt(ctx, debt)
	if err != nil {
		return domain.Algorithm{}, err
	}
	return algo, nil
}

func NewPurchaseUsecase(p domain.PurchaseRepository, u domain.UserRepository, timeout time.Duration) domain.PurchaseRepository {
	return &PurchaseUsecase{
		PurchaseRepo:   p,
		UserRepo:       u,
		contextTimeout: timeout,
	}
}
