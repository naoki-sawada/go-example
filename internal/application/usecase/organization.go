package usecase

import (
	"context"
	"go-example/internal/domain/model"
	"go-example/internal/domain/repository"
)

type OrganizationUseCase interface {
	GetUserList(ctx context.Context, id string) (model.UserList, error)
}

type organizationUseCase struct {
	or repository.OrganizationRepositry
}

func NewOrganizationUseCase(or repository.OrganizationRepositry) OrganizationUseCase {
	return &organizationUseCase{or: or}
}

func (u *organizationUseCase) GetUserList(ctx context.Context, id string) (model.UserList, error) {
	return u.or.GetUserList(ctx, id)
}
