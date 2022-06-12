package repository

import (
	"context"
	"go-example/internal/domain/model"
)

type OrganizationRepositry interface {
	GetByID(ctx context.Context, id string) (*model.Organization, error)
	CreateOrUpdate(ctx context.Context, org *model.Organization) error
	Delete(ctx context.Context, id string) error
	AddUser(ctx context.Context, id, userID string) error
	GetUserList(ctx context.Context, id string) (model.UserList, error)
}
