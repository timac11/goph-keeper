package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/client/cache"
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type Config struct {
	storeDir       string
	privateKeyPath string
	recPath        string
}

type Service struct {
	client Client
	cache  cache.Cache[clientModel.AppSettings]
	config Config
}

type Client interface {
	Login(ctx context.Context, login, password string) (*model.UserLoginResultDto, error)
	Register(ctx context.Context, login, password string) (*model.UserLoginResultDto, error)
	GetList(ctx context.Context, token string) (*[]model.SecretInfoDto, error)
	GetSecret(ctx context.Context, token, id string) (*clientModel.GetSecretDto, error)
	DeleteSecret(ctx context.Context, token, id string) error
	UploadSecret(ctx context.Context, token string, data clientModel.UploadData) error
}
