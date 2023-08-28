/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package main

import (
	"context"

	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http"
	repository "github.com/msantosfelipe/go-bank-transfer/app/repository/db"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/db"
	"github.com/sirupsen/logrus"
)

func main() {
	// log
	configureLog()

	// init db
	ctx := context.Background()
	dbClient := db.InitDb(ctx)
	defer dbClient.Close()
	defer ctx.Done()

	// init dependencies
	accountRepo := repository.NewAccountRepository(dbClient)
	accountUs := usecase.NewAccountUsecase(accountRepo)
	loginRepo := repository.NewLoginRepository(dbClient)
	loginUs := usecase.NewLoginUsecase(loginRepo)
	transferRepo := repository.NewTransferRepository(dbClient)
	transferUs := usecase.NewTransferUsecase(transferRepo)

	// init routers
	http.InitHttpRouters(accountUs, loginUs, transferUs)
}

func configureLog() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
