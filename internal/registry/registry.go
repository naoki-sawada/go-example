package registry

import (
	"go-example/internal/application/usecase"
	"go-example/internal/domain/repository"
	"go-example/internal/domain/service"
	"go-example/internal/infrastructure/persistence/datastore"
	"go-example/internal/interfaces/api/handler"

	"github.com/jmoiron/sqlx"
)

type Registry interface {
	NewUserRepository() repository.UserRepositry
	NewUserService() service.UserService
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
}

type registry struct {
	db *sqlx.DB
}

func NewRegistry(db *sqlx.DB) Registry {
	return &registry{db}
}

func (r registry) NewUserRepository() repository.UserRepositry {
	return datastore.NewUserRepository(r.db)
}

func (r registry) NewUserService() service.UserService {
	ur := r.NewUserRepository()
	return service.NewUserService(ur)
}

func (r registry) NewUserUseCase() usecase.UserUseCase {
	ur := r.NewUserRepository()
	us := r.NewUserService()
	return usecase.NewUserUseCase(ur, us)
}

func (r registry) NewUserHandler() handler.UserHandler {
	us := r.NewUserUseCase()
	return handler.NewUserHandler(us)
}
