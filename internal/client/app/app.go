package app

import (
	"context"

	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type Application struct {
}

type Service interface {
	GetSecretsList(ctx context.Context, args clientModel.ListArgs) (*[]model.SecretInfoDto, error)
	DeleteSecret(ctx context.Context, args clientModel.DeleteArgs) error
	GetSecret(ctx context.Context, args clientModel.GetArgs) (string, error)
	UploadFile(ctx context.Context, args clientModel.UploadFileArgs) error
	UploadAuth(ctx context.Context, args clientModel.UploadAuthArgs) error
	UploadCard(ctx context.Context, args clientModel.UploadCardArgs) error
}
