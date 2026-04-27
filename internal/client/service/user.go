package service

import (
	"context"

	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func (s *Service) Login(ctx context.Context, args clientModel.LoginArgs) (*model.UserLoginResultDto, error) {
	return s.client.Login(ctx, args.Login, args.Password)
}

func (s *Service) Register(ctx context.Context, args clientModel.RegisterArgs) (*model.UserLoginResultDto, error) {
	return s.client.Register(ctx, args.Login, args.Password)
}
