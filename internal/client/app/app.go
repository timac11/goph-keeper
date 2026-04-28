package app

import (
	"context"

	"github.com/timac11/goph-keeper/internal/client/cache"
	"github.com/timac11/goph-keeper/internal/client/cmd"
	"github.com/timac11/goph-keeper/internal/client/config"
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/client/service"
	"github.com/timac11/goph-keeper/internal/client/transport/http"
	"github.com/timac11/goph-keeper/internal/common/logger"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type Application struct {
	service Service
	cache   cache.Cache[clientModel.AppSettings]
}

type Service interface {
	GetSecretsList(ctx context.Context, args clientModel.ListArgs) (*[]model.SecretInfoDto, error)
	DeleteSecret(ctx context.Context, args clientModel.DeleteArgs) error
	GetSecret(ctx context.Context, args clientModel.GetArgs) (string, error)
	UploadFile(ctx context.Context, args clientModel.UploadFileArgs) error
	UploadAuth(ctx context.Context, args clientModel.UploadAuthArgs) error
	UploadCard(ctx context.Context, args clientModel.UploadCardArgs) error
	Register(ctx context.Context, args clientModel.RegisterArgs) error
	Login(ctx context.Context, args clientModel.LoginArgs) error
}

func InitApplication() {
	rootContext := context.Background()
	log := logger.LoggerFromContext(rootContext)

	conf, err := config.InitConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	appClient := http.NewHTTPClient(conf.ServerAddress)
	appCache := cache.NewCache[clientModel.AppSettings](conf.SessionFilePath)
	appService := service.NewService(
		appClient,
		*appCache,
		service.Config{StoreDir: conf.StoreDirPath, PrivateKeyPath: conf.PrivateKeyPath, RecPath: conf.RecPath},
	)

	loginCmd := cmd.BuildLoginCmd(appService.Login)
	registerCmd := cmd.BuildRegisterCmd(appService.Register)

	getSecretListCmd := cmd.BuildListCmd(appService.GetSecretsList)
	deleteSecretCmd := cmd.BuildDeleteCmd(appService.DeleteSecret)
	getSecretCmd := cmd.BuildGetCmd(appService.GetSecret)
	uploadFileCmd := cmd.BuildUploadFileCmd(appService.UploadFile)
	uploadAuthCmd := cmd.BuildUploadAuthCmd(appService.UploadAuth)
	uploadCardCmd := cmd.BuildUploadCardCmd(appService.UploadCard)

	cmd.RootCmd.AddCommand(
		loginCmd,
		registerCmd,
		getSecretListCmd,
		deleteSecretCmd,
		getSecretCmd,
		uploadFileCmd,
		uploadAuthCmd,
		uploadCardCmd,
	)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
