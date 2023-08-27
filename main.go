/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase"
	"github.com/msantosfelipe/go-bank-transfer/config"
	"github.com/sirupsen/logrus"
)

func main() {
	// log
	configureLog()

	// init db
	ctx := context.Background()
	dbClient, err := pgxpool.Connect(ctx, getDbUri())
	if err != nil {
		panic(err)
	}
	defer dbClient.Close()

	// init dependencies
	accountRepo := db.NewAccountRepository(dbClient)
	accountUs := usecase.NewAccountUsecase(accountRepo)
	loginRepo := db.NewLoginRepository(dbClient)
	loginUs := usecase.NewLoginUsecase(loginRepo)
	transferRepo := db.NewTransferRepository(dbClient)
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

func getDbUri() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.ENV.DbUser,
		config.ENV.DbPass,
		config.ENV.DbHost,
		config.ENV.DbPort,
		config.ENV.DbName,
	)
}
