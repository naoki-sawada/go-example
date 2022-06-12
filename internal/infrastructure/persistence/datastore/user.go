package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-example/internal/domain/model"
	"go-example/internal/domain/repository"
	e "go-example/internal/utils/errors"

	"github.com/jmoiron/sqlx"
)

type UserRepositry struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepositry {
	return &UserRepositry{db: db}
}

func (r UserRepositry) GetByID(ctx context.Context, id string) (*model.User, error) {
	const query = `
		SELECT "id", "first_name", "last_name", "email", "birthdate"
		FROM "users"
		WHERE "id" = $1
	`
	var u model.User
	if err := r.db.GetContext(ctx, &u, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.NewErrNotFound(err)
		}
		return nil, err
	}
	return &u, nil
}

func (r UserRepositry) CreateOrUpdate(ctx context.Context, user *model.User) error {
	const query = `
		INSERT INTO "users" ("id", "first_name", "last_name", "email", "birthdate")
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT ("id")
    DO UPDATE SET "first_name" = $2, "last_name" = $3, "email" = $4, "birthdate" = $5 
	`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.Birthdate)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepositry) Delete(ctx context.Context, id string) error {
	const query = `
		DELETE FROM "users"
		WHERE "id" = $1
	`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return e.NewErrNotFound(fmt.Errorf("no rows affected"))
	}
	return nil
}
