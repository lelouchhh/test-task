package domain

import "context"

type Algorithm struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Task     string  `json:"task"`
	Solution string  `json:"solution" db:"solution"`
	Price    float64 `json:"price" db:"price"`
}

type AlgorithmUsecase interface {
	GetAlgorithms(ctx context.Context) ([]Algorithm, error)
	GetSolution(ctx context.Context, id string) (Algorithm, error)
}

type AlgorithmRepository interface {
	GetAlgorithms(ctx context.Context) ([]Algorithm, error)
	GetSolution(ctx context.Context, id string) (Algorithm, error)
}
