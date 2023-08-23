/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/msantosfelipe/go-bank-transfer/app/delivery/http"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase"
	"github.com/msantosfelipe/go-bank-transfer/config"
)

func main() {
	// init dependencies
	accountUs := usecase.NewAccountUsecase()

	// init routers
	engine := gin.New()
	apiRouter := engine.Group("/go-bank-transfer")
	http.NewAccountRouter(apiRouter, accountUs)

	// serve
	engine.Run(":" + config.ENV.ApiPort)

}
