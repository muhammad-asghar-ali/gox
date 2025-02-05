package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	UserActions interface {
		Create(ctx context.Context, req *entities.CreateUserParams) *entities.User
	}

	UserService struct{}
)

func NewUserService() UserActions {
	return &UserService{}
}

func (us *UserService) Create(ctx context.Context, req *entities.CreateUserParams) *entities.User {
	return nil
}
