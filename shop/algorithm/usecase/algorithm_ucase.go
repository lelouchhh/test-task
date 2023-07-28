package usecase

import (
	"context"
	"shop/domain"
	"time"
)

type AlgorithmUsecase struct {
	AlgoRepo       domain.AlgorithmRepository
	contextTimeout time.Duration
}

func (a AlgorithmUsecase) GetAlgorithms(ctx context.Context) ([]domain.Algorithm, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	algorithms, err := a.AlgoRepo.GetAlgorithms(ctx)
	if err != nil {
		return []domain.Algorithm{}, err
	}
	return algorithms, nil
}

func NewUserUsecase(u domain.AlgorithmRepository, timeout time.Duration) domain.AlgorithmUsecase {
	return &AlgorithmUsecase{
		AlgoRepo:       u,
		contextTimeout: timeout,
	}
}
