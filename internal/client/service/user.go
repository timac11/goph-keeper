package service

import (
	"context"

	clientModel "github.com/timac11/goph-keeper/internal/client/model"
)

func (s *Service) Login(ctx context.Context, args clientModel.LoginArgs) error {
	res, err := s.client.Login(ctx, args.Login, args.Password)
	if err != nil {
		return err
	}

	appSettings := clientModel.AppSettings{Token: res.Token}
	return s.cache.Store(ctx, appSettings)
}

func (s *Service) Register(ctx context.Context, args clientModel.RegisterArgs) error {
	res, err := s.client.Register(ctx, args.Login, args.Password)
	if err != nil {
		return err
	}

	appSettings := clientModel.AppSettings{Token: res.Token}
	return s.cache.Store(ctx, appSettings)
}
