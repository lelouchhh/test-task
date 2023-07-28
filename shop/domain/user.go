package domain

import "context"

type User struct {
	Id         string  `json:"id"`
	FirstName  string  `json:"first_name" validate:"required,gt=2"`
	SecondName string  `json:"second_name" validate:"required,gt=2"`
	LastName   string  `json:"last_name" validate:"required,gt=2"`
	Email      string  `json:"email" validate:"required,email"`
	Debt       float64 `json:"debt" validate:"omitempty"`
}
type Debt struct {
	Id   string  `json:"id" validate:"required"`
	Debt float32 `json:"debt" validate:"required"`
}

type UserUsecase interface {
	CreateUser(ctx context.Context, user User) error
	IncreaseDebt(ctx context.Context, amount Debt) error
	DecreaseDebt(ctx context.Context, amount Debt) error
}
type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	IncreaseDebt(ctx context.Context, amount Debt) error
	DecreaseDebt(ctx context.Context, amount Debt) error
}
