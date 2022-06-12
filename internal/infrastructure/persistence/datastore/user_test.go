package datastore

import (
	"context"
	"go-example/internal/domain/model"
	e "go-example/internal/utils/errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdate(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	u := &model.User{
		ID:        uuid.NewString(),
		FirstName: "world",
		LastName:  "hello",
		Email:     "hello@world",
		Birthdate: time.Now(),
	}

	// create
	r := NewUserRepository(db)
	err = r.CreateOrUpdate(ctx, u)
	assert.NoError(t, err)

	res, err := r.GetByID(ctx, u.ID)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, res.ID)
	assert.Equal(t, u.FirstName, res.FirstName)
	assert.Equal(t, u.LastName, res.LastName)
	assert.Equal(t, u.Email, res.Email)
	assert.WithinDuration(t, u.Birthdate, res.Birthdate, time.Hour*24)

	// update
	uu := &model.User{
		ID:        u.ID,
		FirstName: "updated",
		LastName:  "name",
		Email:     "updated@name",
		Birthdate: time.Now().AddDate(10, 0, 0),
	}

	err = r.CreateOrUpdate(ctx, uu)
	assert.NoError(t, err)

	res, err = r.GetByID(ctx, u.ID)
	assert.NoError(t, err)
	assert.Equal(t, uu.ID, res.ID)
	assert.Equal(t, uu.FirstName, res.FirstName)
	assert.Equal(t, uu.LastName, res.LastName)
	assert.Equal(t, uu.Email, res.Email)
	assert.WithinDuration(t, uu.Birthdate, res.Birthdate, time.Hour*24)
}

func TestGetByID(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	id := uuid.NewString()

	r := NewUserRepository(db)
	u, err := r.GetByID(ctx, id)

	var en *e.ErrNotFound
	assert.ErrorAs(t, err, &en)
	assert.Nil(t, u)
}

func TestDeleteUserNotFound(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	id := uuid.NewString()

	r := NewUserRepository(db)
	err = r.Delete(ctx, id)

	var en *e.ErrNotFound
	assert.ErrorAs(t, err, &en)
}

func TestDelete(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	u := &model.User{
		ID:        uuid.NewString(),
		FirstName: "world",
		LastName:  "hello",
		Email:     "hello@world",
		Birthdate: time.Now(),
	}

	r := NewUserRepository(db)
	err = r.CreateOrUpdate(ctx, u)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Delete(ctx, u.ID)
	assert.NoError(t, err)
}
