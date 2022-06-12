package service

import (
	"context"
	"errors"
	"go-example/internal/domain/model"
	"go-example/internal/domain/repository"
	e "go-example/internal/utils/errors"
)

type UserService interface {
	Exists(ctx context.Context, user *model.User) (bool, error)
}

type userService struct {
	ur repository.UserRepositry
}

func NewUserService(ur repository.UserRepositry) UserService {
	return &userService{ur: ur}
}

func (s userService) Exists(ctx context.Context, user *model.User) (bool, error) {
	user, err := s.ur.GetByID(ctx, user.ID)
	if err != nil {
		var en *e.ErrNotFound
		if errors.As(err, &en) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
