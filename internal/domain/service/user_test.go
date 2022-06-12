package service

import (
	"context"
	"fmt"
	"go-example/internal/domain/model"
	"go-example/internal/domain/repository"
	e "go-example/internal/utils/errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type config struct {
	getByID func(ctx context.Context, id string) (*model.User, error)
}

type userRepository struct {
	c *config
}

func newUserRepository(c *config) repository.UserRepositry {
	return &userRepository{c: c}
}

func (r userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	return r.c.getByID(ctx, id)
}

func (r userRepository) CreateOrUpdate(ctx context.Context, user *model.User) error {
	return nil
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func TestUserExists(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		config   *config
		user     *model.User
		expected bool
	}{
		{
			config: &config{
				getByID: func(ctx context.Context, id string) (*model.User, error) {
					return &model.User{ID: id}, nil
				},
			},
			user:     &model.User{ID: uuid.NewString()},
			expected: true,
		},
		{
			config: &config{
				getByID: func(ctx context.Context, id string) (*model.User, error) {
					return nil, e.NewErrNotFound(fmt.Errorf("not found"))
				},
			},
			user:     &model.User{ID: uuid.NewString()},
			expected: false,
		},
	}

	for i, tc := range cases {
		ur := newUserRepository(tc.config)
		s := NewUserService(ur)
		actual, err := s.Exists(ctx, tc.user)
		assert.NoError(t, err, fmt.Sprintf("case %d has error", i))
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("case %d failed", i))
	}
}
