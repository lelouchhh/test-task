package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"shop/domain"
)

type psqlUserRepository struct {
	db *sqlx.DB
}

func (p psqlUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := p.doesUserExist(ctx, user)
	if err != nil {
		return err
	}
	query := `insert into users (first_name, second_name, last_name, email) values ($1, $2, $3, $4)
`
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, user.FirstName, user.SecondName, user.LastName, user.Email)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (p psqlUserRepository) IncreaseDebt(ctx context.Context, amount domain.Debt) error {
	// Begin the transaction
	tx, err := p.db.Begin()
	if err != nil {
		return domain.ErrInternalServerError
	}

	// Lock the row for the user with the given ID
	_, err = tx.ExecContext(ctx, "SELECT debt FROM users WHERE id = $1 FOR UPDATE", amount.Id)
	if err != nil {
		tx.Rollback()
		if err != nil {
			return err
		}
		return domain.ErrInternalServerError
	}

	// Update the user's balance safely within the transaction
	_, err = tx.ExecContext(ctx, "UPDATE users SET debt = debt + $1 WHERE id = $2", amount.Debt, amount.Id)
	if err != nil {
		tx.Rollback()
		is := isConstraintViolation(err)
		if is {
			return domain.ErrUnprocessableEntity
		}
		return domain.ErrInternalServerError
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.ErrInternalServerError
	}

	return nil
}

func (p psqlUserRepository) DecreaseDebt(ctx context.Context, amount domain.Debt) error {
	// Begin the transaction
	tx, err := p.db.Begin()
	if err != nil {
		return domain.ErrInternalServerError
	}

	// Lock the row for the user with the given ID
	_, err = tx.ExecContext(ctx, "SELECT debt FROM users WHERE id = $1 FOR UPDATE", amount.Id)
	if err != nil {
		tx.Rollback()
		if err != nil {
			return err
		}
		return domain.ErrInternalServerError
	}

	// Update the user's balance safely within the transaction
	_, err = tx.ExecContext(ctx, "UPDATE users SET debt = debt - $1 WHERE id = $2", amount.Debt, amount.Id)
	if err != nil {
		tx.Rollback()
		is := isConstraintViolation(err)
		if is {
			return domain.ErrUnprocessableEntity
		}
		return domain.ErrInternalServerError
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.ErrInternalServerError
	}

	return nil
}

func (a *psqlUserRepository) doesUserExist(ctx context.Context, user domain.User) error {
	var doesExist bool
	err := a.db.GetContext(ctx, &doesExist, "select 1 from users where email = $1", user.Email)

	if err == nil {
		return domain.ErrConflict
	}
	return nil
}
func isConstraintViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code.Name() == "check_violation" || pqErr.Code.Name() == "not_null_violation" || pqErr.Code.Name() == "unique_violation"
	}
	return false
}
func NewUserRepository(conn *sqlx.DB) domain.UserRepository {
	return &psqlUserRepository{conn}
}
