//go:build wireinject

import (
	"backend-go/internal/controller"
	"backend-go/internal/repo"
	"backend-go/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	// declare dependencies in order
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}