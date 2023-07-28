package usecase

import (
	"context"
	"stock-service/domain"
	"time"
)

type algorithmUsecase struct {
	AlgoRepo       domain.AlgorithmRepository
	contextTimeout time.Duration
}

func (a algorithmUsecase) GetAlgorithms(ctx context.Context) ([]domain.Algorithm, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	algorithms, err := a.AlgoRepo.GetAlgorithms(ctx)
	if err != nil {
		return []domain.Algorithm{}, err
	}
	return algorithms, nil
}

func (a algorithmUsecase) GetSolution(ctx context.Context, id string) (domain.Algorithm, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	solution, err := a.AlgoRepo.GetSolution(ctx, id)
	if err != nil {
		return domain.Algorithm{}, err
	}
	return solution, nil
}

func NewAlgorithmUsecase(a domain.AlgorithmRepository, timeout time.Duration) domain.AlgorithmUsecase {
	return &algorithmUsecase{
		AlgoRepo:       a,
		contextTimeout: timeout,
	}
}
