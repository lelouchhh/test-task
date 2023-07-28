package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"stock-service/domain"
)

type psqlAlgorithmRepository struct {
	db *sqlx.DB
}

func (p psqlAlgorithmRepository) GetAlgorithms(ctx context.Context) ([]domain.Algorithm, error) {
	var algs []domain.Algorithm
	err := p.db.SelectContext(ctx, &algs, "select id, name, price, task from algorithm")
	if err != nil {
		return nil, domain.ErrInternalServerError
	}
	return algs, nil
}

func (p psqlAlgorithmRepository) GetSolution(ctx context.Context, id string) (domain.Algorithm, error) {
	var solution domain.Algorithm
	err := p.db.GetContext(ctx, &solution, "select solution, price from algorithm where id = $1", id)
	if err != nil {
		return domain.Algorithm{}, domain.ErrInternalServerError
	}
	return solution, nil
}

func NewAlgorithmRepository(conn *sqlx.DB) domain.AlgorithmRepository {
	return &psqlAlgorithmRepository{conn}
}
