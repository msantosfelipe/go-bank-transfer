/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/logger"
)

type accountHandler struct {
	accountUs domain.AccountUsecase
}

// NewAccountRouter handle REST requests
func NewAccountRouter(router *gin.RouterGroup, accountUs domain.AccountUsecase) {
	handler := accountHandler{
		accountUs: accountUs,
	}

	router.Use(
		logger.LogMiddleware,
	)

	router.POST("/accounts", handler.createAccount)
}

func (handler *accountHandler) createAccount(context *gin.Context) {
	var body domain.Account
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusPreconditionFailed, domain.BuildResponseFromError(err))
		return
	}

	context.Status(http.StatusOK)
}
