/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase"
	"github.com/msantosfelipe/go-bank-transfer/config"
)

func main() {
	// init db
	ctx := context.Background()
	dbClient, err := pgxpool.Connect(ctx, config.ENV.DbUri)
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()

	// init dependencies
	accountRepo := db.NewAccountRepository(dbClient)
	accountUs := usecase.NewAccountUsecase(accountRepo)
	loginRepo := db.NewLoginRepository(dbClient)
	loginUs := usecase.NewLoginUsecase(loginRepo)

	// init routers
	http.InitHttpRouters(accountUs, loginUs)
}
