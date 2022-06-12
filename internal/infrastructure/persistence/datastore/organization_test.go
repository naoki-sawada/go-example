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

func TestOrganiztionCreateOrUpdate(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	o := &model.Organization{
		ID:   uuid.NewString(),
		Name: "Test Inc.",
	}

	// create
	r := NewOrganizationRepository(db)
	err = r.CreateOrUpdate(ctx, o)
	assert.NoError(t, err)

	res, err := r.GetByID(ctx, o.ID)
	assert.NoError(t, err)
	assert.Equal(t, o.ID, res.ID)
	assert.Equal(t, o.Name, res.Name)

	// update
	uo := &model.Organization{
		ID:   o.ID,
		Name: "New Test Inc.",
	}

	err = r.CreateOrUpdate(ctx, uo)
	assert.NoError(t, err)

	res, err = r.GetByID(ctx, uo.ID)
	assert.NoError(t, err)
	assert.Equal(t, uo.ID, res.ID)
	assert.Equal(t, uo.Name, res.Name)
}

func TestOrganizationGetByID(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	id := uuid.NewString()

	r := NewOrganizationRepository(db)
	u, err := r.GetByID(ctx, id)

	var en *e.ErrNotFound
	assert.ErrorAs(t, err, &en)
	assert.Nil(t, u)
}

func TestDeleteOrganizationNotFound(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	id := uuid.NewString()

	r := NewOrganizationRepository(db)
	err = r.Delete(ctx, id)

	var en *e.ErrNotFound
	assert.ErrorAs(t, err, &en)
}

func TestOrganizationDelete(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	o := &model.Organization{
		ID:   uuid.NewString(),
		Name: "Test Ltd.",
	}

	r := NewOrganizationRepository(db)
	err = r.CreateOrUpdate(ctx, o)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Delete(ctx, o.ID)
	assert.NoError(t, err)
}

func TestOrganizationAddUser(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	o := &model.Organization{
		ID:   uuid.NewString(),
		Name: "Test LLC",
	}
	or := NewOrganizationRepository(db)
	err = or.CreateOrUpdate(ctx, o)
	if err != nil {
		t.Fatal(err)
	}
	u := &model.User{
		ID:        uuid.NewString(),
		FirstName: "Alexandrea",
		LastName:  "Armstrong",
		Email:     "aa@test",
		Birthdate: time.Now(),
	}
	ur := NewUserRepository(db)
	err = ur.CreateOrUpdate(ctx, u)
	if err != nil {
		t.Fatal(err)
	}

	err = or.AddUser(ctx, o.ID, u.ID)
	assert.NoError(t, err)

	ul, err := or.GetUserList(ctx, o.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ul))
	assert.Equal(t, u.ID, ul[0].ID)
}

func TestOrganizationGetUserList(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()

	orgID := uuid.NewString()

	or := NewOrganizationRepository(db)

	ul, err := or.GetUserList(ctx, orgID)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(ul))
}
