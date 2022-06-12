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

type OrganizationRepositry struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) repository.OrganizationRepositry {
	return &OrganizationRepositry{db: db}
}

func (r OrganizationRepositry) GetByID(ctx context.Context, id string) (*model.Organization, error) {
	const query = `
		SELECT "id", "name"
		FROM "organizations"
		WHERE "id" = $1
	`
	var o model.Organization
	if err := r.db.GetContext(ctx, &o, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.NewErrNotFound(err)
		}
		return nil, err
	}
	return &o, nil
}

func (r OrganizationRepositry) CreateOrUpdate(ctx context.Context, org *model.Organization) error {
	const query = `
		INSERT INTO "organizations" ("id", "name")
		VALUES ($1, $2)
		ON CONFLICT ("id")
    DO UPDATE SET "name" = $2
	`
	_, err := r.db.ExecContext(ctx, query, org.ID, org.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r OrganizationRepositry) Delete(ctx context.Context, id string) error {
	const query = `
		DELETE FROM "organizations"
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

func (r OrganizationRepositry) GetUserList(ctx context.Context, id string) (model.UserList, error) {
	const query = `
		SELECT "u"."id", "u"."first_name", "u"."last_name", "u"."email", "u"."birthdate"
		FROM "users" "u"
		INNER JOIN "organization_user" "ou" ON "u"."id" = "ou"."user_id"
		WHERE "ou"."organization_id" = $1
	`
	var ul model.UserList
	if err := r.db.SelectContext(ctx, &ul, query, id); err != nil {
		return nil, err
	}
	return ul, nil
}

func (r OrganizationRepositry) AddUser(ctx context.Context, id, userID string) error {
	const query = `
		INSERT INTO "organization_user" ("organization_id", "user_id")
		VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}
	return nil
}
