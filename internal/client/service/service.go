package service

import (
	"github.com/timac11/goph-keeper/internal/client/cache"
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type Service struct {
	client Client
	cache  cache.Cache[clientModel.AppSettings]
}

type Client interface {
	Login(login, password string) (*model.UserLoginResultDto, error)
	Register(login, password string) (*model.UserLoginResultDto, error)
	GetList(token string) (*[]model.SecretInfoDto, error)
	GetSecret(token, id string) (*model.Secret, error)
	DeleteSecret(token, id string) error
	UploadSecret(token string, data clientModel.UploadData) error
}
