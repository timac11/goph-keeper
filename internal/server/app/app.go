package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/timac11/goph-keeper/internal/common/logger"
	"github.com/timac11/goph-keeper/internal/server/auth"
	"github.com/timac11/goph-keeper/internal/server/config"
	"github.com/timac11/goph-keeper/internal/server/repository"
	"github.com/timac11/goph-keeper/internal/server/service"
	"github.com/timac11/goph-keeper/internal/server/transport/http"
	"github.com/timac11/goph-keeper/internal/server/transport/http/handler"
	"golang.org/x/sync/errgroup"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

func InitApplication() {
	conf := config.InitConfig()

	dbClient, err := repository.NewPgClient(conf.DatabaseURI)

	if err != nil {
		log.Fatal(err)
	}

	appService := service.NewService(dbClient)
	jwtControl := auth.JWTControl{
		TokenExp: (time.Duration(conf.JWTExpMinutes * uint(time.Minute))),
		Secret:   conf.JWTSecret,
	}

	server := initHttpServer(conf.Address, appService, &jwtControl)

	g, appCtx := errgroup.WithContext(context.Background())
	defer appCtx.Err()

	log := logger.LoggerFromContext(appCtx)

	// run server
	g.Go(func() error {
		log.Info("start server")
		return server.Start()
	})

	// graceful shutdown
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		<-quitChan
		log.Info("graceful shutdown signal")

		if err = server.Stop(appCtx); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err.Error())
	}
}

func initHttpServer(address string, appService *service.Service, jwtControl *auth.JWTControl) Server {
	handler := handler.NewHandler(appService, jwtControl)
	router := http.Init(*handler, jwtControl)

	return http.NewServer(address, router)
}
