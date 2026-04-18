package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/common/model"
)

type Service struct {
	repository Repository
}

type Repository interface {
	// user api
	CreateUser(ctx context.Context, value *model.UserLoginDto) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	// secret api
	CreateSecret(ctx context.Context, secret *model.SecretCreateDto, userId string) (*model.SecretInfoDto, error)
	GetSecret(ctx context.Context, id string) (*model.Secret, error)
	DeleteSecret(ctx context.Context, id string) error
	UpdateSecret(ctx context.Context, secret *model.Secret) (*model.Secret, error)
	GetSecretList(ctx context.Context, id string) (*[]model.SecretInfoDto, error)
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
