package service

import (
	"context"

	"github.com/timac11/goph-keeper/internal/common/model"
)

func (service *Service) CreateSecret(ctx context.Context, secret *model.SecretCreateDto, userId string) (*model.SecretInfoDto, error) {
	return service.repository.CreateSecret(ctx, secret, userId)
}

func (service *Service) GetSecret(ctx context.Context, id string) (*model.Secret, error) {
	return service.repository.GetSecret(ctx, id)
}

func (service *Service) UpdateSecret(ctx context.Context, secret *model.Secret) (*model.Secret, error) {
	return service.repository.UpdateSecret(ctx, secret)
}
func (service *Service) DeleteSecret(ctx context.Context, id string) error {
	return service.repository.DeleteSecret(ctx, id)
}
