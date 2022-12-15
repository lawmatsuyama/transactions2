package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lawmatsuyama/pismo-transactions/docs"
	"github.com/lawmatsuyama/pismo-transactions/infra/apimanager"
	"github.com/lawmatsuyama/pismo-transactions/infra/repository"
	log "github.com/sirupsen/logrus"
)

var (
	serviceName = "accounts-service"
)

func init() {
	LoadEnv()
	LoggerSetup()
}

// @title Swagger Accounts API
// @version 2.0
// @description API to create,list account and transactions of user.
// @termsOfService http://swagger.io/terms/

// @contact.name Lawrence Matsuyama
// @contact.email law.matsuyama@outlook.com

// @host localhost:8080
// @BasePath /
func main() {
	ctx, cancel := start()
	defer shutdown(ctx, cancel)
	waitSignal()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	log.WithField("service", serviceName).Info("Starting service")
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	StartDependencies(ctxWithCancel)
	log.WithField("service", serviceName).Info("Service is ready")
	return ctxWithCancel, cancel
}

func shutdown(ctx context.Context, cancel context.CancelFunc) {
	cancel()
	err := repository.CloseDB(ctx)
	if err != nil {
		log.Warn("Failed to close database")
	}
	apimanager.ShutdownAPI()
	log.WithField("service", serviceName).Info("Service finished")
}

func waitSignal() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.WithField("service", serviceName).Infof("Signal received [%v] canceling everything", s)
}
