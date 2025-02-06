package services

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/db"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	UserActions interface {
		Create(ctx context.Context, req *entities.CreateUserParams) (*entities.User, error)
		Login(ctx context.Context, req *common.LoginRequest) (*common.LoginResponse, error)
		FindUserByEmail(ctx context.Context, email string) (*entities.User, error)
	}

	UserService struct{}
)

func NewUserService() UserActions {
	return &UserService{}
}

func (us *UserService) Create(ctx context.Context, req *entities.CreateUserParams) (*entities.User, error) {
	req.Password = common.HashPassword(req.Password)
	user, err := db.Queries().CreateUser(ctx, *req)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) Login(ctx context.Context, req *common.LoginRequest) (*common.LoginResponse, error) {
	user, err := db.Queries().FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := common.ComparePassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	access, err := common.AccessToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	refresh, err := common.RefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	res := &common.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return res, nil
}

func (us *UserService) FindUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := db.Queries().FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
