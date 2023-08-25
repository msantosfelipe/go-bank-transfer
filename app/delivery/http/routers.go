/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http/middleware"
	"github.com/msantosfelipe/go-bank-transfer/config"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	docs "github.com/msantosfelipe/go-bank-transfer/infrastructure/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitHttpRouters(
	accountUs domain.AccountUsecase,
	loginUs domain.LoginUsecase,
	transferUs domain.TransferUsecase,
) {
	engine := gin.New()
	apiRouter := engine.Group(config.ENV.ApiBasePath)

	docs.SwaggerInfo.BasePath = config.ENV.ApiBasePath
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// public routes
	apiRouter.Use(
		middleware.LogMiddleware,
	)

	NewAccountHandler(apiRouter, accountUs)
	NewLoginHandler(apiRouter, loginUs)

	// protected routes
	apiRouter.Use(
		middleware.JwthMiddleware,
	)

	NewTransferHandler(apiRouter, transferUs)

	// run app
	engine.Run(":" + config.ENV.ApiPort)
}
