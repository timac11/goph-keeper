package service

import (
	"github.com/timac11/goph-keeper/internal/client/cache"
	"github.com/timac11/goph-keeper/internal/client/model"
)

type Service struct {
	client Client
	cache  cache.Cache[model.AppSettings]
}

type Client struct{}
