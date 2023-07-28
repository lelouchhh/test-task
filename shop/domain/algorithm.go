package domain

import "context"

type Algorithm struct {
	Id       string
	Name     string
	Task     string
	Solution string
	Price    float32
}

type AlgorithmUsecase interface {
	GetAlgorithms(ctx context.Context) ([]Algorithm, error)
}

type AlgorithmRepository interface {
	GetAlgorithms(ctx context.Context) ([]Algorithm, error)
}
