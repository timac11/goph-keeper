package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/common/util"
	"github.com/timac11/goph-keeper/internal/server/errors"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func (service *Service) Register(ctx context.Context, value *model.UserLoginDto) (*model.User, error) {

	password, err := util.HashPassword(value.Password)

	if err != nil {
		return nil, err
	}

	userModel, err := service.Repository.CreateUser(ctx, &model.UserLoginDto{Login: value.Login, Password: password})

	if err != nil {
		return nil, err
	}

	return userModel, nil

}

func (service *Service) Login(ctx context.Context, value *model.UserLoginDto) (*model.User, error) {
	userModel, err := service.Repository.GetUserByLogin(ctx, value.Login)

	if err != nil {
		return nil, err
	}

	passwordCorrect := util.CheckPasswordHash(value.Password, userModel.Password)

	if !passwordCorrect {
		return nil, errors.NewInvalidPasswordError(value.Password)
	}

	return userModel, nil
}
