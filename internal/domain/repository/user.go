package repository

import (
	"context"
	"go-example/internal/domain/model"
)

type UserRepositry interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	CreateOrUpdate(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
