package usecase

import (
	"context"
	"go-example/internal/domain/model"
	"go-example/internal/domain/repository"
	"go-example/internal/domain/service"
	e "go-example/internal/utils/errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, firstName, lastName, email string, birthdate time.Time) (*model.User, error)
	UpdateUser(ctx context.Context, id string, firstName, lastName, email *string, birthdate *time.Time) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userUseCase struct {
	ur repository.UserRepositry
	us service.UserService
}

func NewUserUseCase(r repository.UserRepositry, s service.UserService) UserUseCase {
	return &userUseCase{ur: r, us: s}
}

func (c userUseCase) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return c.ur.GetByID(ctx, id)
}

func (c userUseCase) CreateUser(ctx context.Context, firstName, lastName, email string, birthdate time.Time) (*model.User, error) {
	u := &model.User{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Birthdate: birthdate,
	}
	if err := validator.New().Struct(u); err != nil {
		return nil, e.NewErrInvalidData(err)
	}
	if err := c.ur.CreateOrUpdate(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (c userUseCase) UpdateUser(ctx context.Context, id string, firstName, lastName, email *string, birthdate *time.Time) (*model.User, error) {
	u, err := c.ur.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if firstName != nil {
		u.FirstName = *firstName
	}
	if lastName != nil {
		u.LastName = *lastName
	}
	if email != nil {
		u.Email = *email
	}
	if birthdate != nil {
		u.Birthdate = *birthdate
	}
	if err := validator.New().Struct(u); err != nil {
		return nil, e.NewErrInvalidData(err)
	}
	if err := c.ur.CreateOrUpdate(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (c userUseCase) DeleteUser(ctx context.Context, id string) error {
	return c.ur.Delete(ctx, id)
}
