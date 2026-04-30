package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/common/model"
	"github.com/timac11/goph-keeper/internal/server/auth"
)

func (service *Service) CreateSecret(ctx context.Context, secret *model.SecretCreateDto) (*model.SecretInfoDto, error) {
	payload, err := auth.AuthPayloadFromContext(ctx)

	if err != nil {
		return nil, err
	}

	return service.repository.CreateSecret(ctx, secret, payload.UserID)
}

func (service *Service) GetSecret(ctx context.Context, id string) (*model.Secret, error) {
	payload, err := auth.AuthPayloadFromContext(ctx)

	if err != nil {
		return nil, err
	}

	return service.repository.GetSecret(ctx, id, payload.UserID)
}

func (service *Service) DeleteSecret(ctx context.Context, id string) error {
	payload, err := auth.AuthPayloadFromContext(ctx)

	if err != nil {
		return err
	}

	return service.repository.DeleteSecret(ctx, id, payload.UserID)
}

func (service *Service) GetSecretsList(ctx context.Context) (*[]model.SecretInfoDto, error) {
	payload, err := auth.AuthPayloadFromContext(ctx)

	if err != nil {
		return nil, err
	}

	return service.repository.GetSecretsList(ctx, payload.UserID)
}
