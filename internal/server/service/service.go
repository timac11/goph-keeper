package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/common/model"
)

type Service struct {
	Repository
}

type Repository interface {
	CreateUser(ctx context.Context, value *model.UserLoginDto) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}
