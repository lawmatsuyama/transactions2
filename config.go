package main

import (
	"context"

	"github.com/joho/godotenv"
	joonix "github.com/joonix/log"
	"github.com/lawmatsuyama/pismo-transactions/infra/apimanager"
	"github.com/lawmatsuyama/pismo-transactions/infra/repository"
	"github.com/lawmatsuyama/pismo-transactions/usecases"

	log "github.com/sirupsen/logrus"
)

// LoadEnv load all enviroment variables from file if it is not already loaded
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Error("couldnt load .env")
	}
}

// LoggerSetup setup log format
func LoggerSetup() {
	log.SetFormatter(joonix.NewFormatter(joonix.PrettyPrintFormat, joonix.DefaultFormat))
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)
}

// StartDependencies it will start and inject dependencies into api
func StartDependencies(ctxWithCancel context.Context) {
	dbCli := repository.Start(context.Background())

	transactionRepository := repository.NewTransactionRepository(dbCli)
	accountRepository := repository.NewAccountRepository(dbCli)

	transactionUseCase := usecases.NewTransactionUseCase(transactionRepository, accountRepository)
	accountUseCase := usecases.NewAccountUseCase(accountRepository)

	transactionAPI := apimanager.NewTransactionAPI(transactionUseCase)
	accountAPI := apimanager.NewAccountAPI(accountUseCase)

	handler := apimanager.NewHandler(accountAPI, transactionAPI)
	apimanager.StartAPI(ctxWithCancel, handler, "8080", "pismo-transactions")

}
