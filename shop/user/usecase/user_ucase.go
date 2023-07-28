package usecase

import (
	"context"
	"shop/domain"
	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func (u UserUsecase) CreateUser(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) IncreaseDebt(ctx context.Context, amount domain.Debt) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.UserRepo.IncreaseDebt(ctx, amount)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) DecreaseDebt(ctx context.Context, amount domain.Debt) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	err := u.UserRepo.DecreaseDebt(ctx, amount)
	if err != nil {
		return err
	}
	return nil
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserRepository {
	return &UserUsecase{
		UserRepo:       u,
		contextTimeout: timeout,
	}
}
