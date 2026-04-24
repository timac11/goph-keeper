package service

import (
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func (s *Service) Login(args clientModel.LoginArgs) (*model.UserLoginResultDto, error) {
	return s.client.Login(args.Login, args.Password)
}

func (s *Service) Register(args clientModel.RegisterArgs) (*model.UserLoginResultDto, error) {
	return s.client.Register(args.Login, args.Password)
}
