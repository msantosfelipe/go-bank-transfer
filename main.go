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

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase"
	"github.com/msantosfelipe/go-bank-transfer/config"
	docs "github.com/msantosfelipe/go-bank-transfer/infrastructure/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const basePath = "/go-bank-transfer"

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
	engine := gin.New()
	apiRouter := engine.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	http.NewAccountRouter(apiRouter, accountUs)
	http.NewLoginRouter(apiRouter, loginUs)

	// serve
	engine.Run(":" + config.ENV.ApiPort)
}
