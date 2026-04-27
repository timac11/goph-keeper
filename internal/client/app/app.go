package app

import (
	"context"

	"github.com/timac11/goph-keeper/internal/client/cache"
	"github.com/timac11/goph-keeper/internal/client/cmd"
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type Application struct {
	service Service
	cache  cache.Cache[clientModel.AppSettings]
}

type Service interface {
	GetSecretsList(ctx context.Context, args clientModel.ListArgs) (*[]model.SecretInfoDto, error)
	DeleteSecret(ctx context.Context, args clientModel.DeleteArgs) error
	GetSecret(ctx context.Context, args clientModel.GetArgs) (string, error)
	UploadFile(ctx context.Context, args clientModel.UploadFileArgs) error
	UploadAuth(ctx context.Context, args clientModel.UploadAuthArgs) error
	UploadCard(ctx context.Context, args clientModel.UploadCardArgs) error
	Register(ctx context.Context, args clientModel.RegisterArgs) (*model.UserLoginResultDto, error)
	Login(ctx context.Context, args clientModel.LoginArgs) (*model.UserLoginResultDto, error)
}

func NewApplication(service Service)*Application {
	return &Application{ service: service }
}

func (app *Application) Login(ctx context.Context, args clientModel.LoginArgs) error {
	res, err := app.service.Login(ctx, args)
	if err != nil {
		return err
	}

	appSettings := clientModel.AppSettings{ Token: res.Token }
	err = app.cache.Store(appSettings)

	return nil
}

func (app *Application) Register(ctx context.Context, args clientModel.RegisterArgs) error {
	res, err := app.service.Register(ctx, args)
	if err != nil {
		return err
	}

	appSettings := clientModel.AppSettings{ Token: res.Token }
	err = app.cache.Store(appSettings)

	return nil
}

func (app *Application) Run() {
	loginCmd := cmd.BuildLoginCmd(app.Login)
	registerCmd := cmd.BuildRegisterCmd(app.Register)

	cmd.RootCmd.AddCommand(loginCmd)
	cmd.RootCmd.AddCommand(registerCmd)
}
