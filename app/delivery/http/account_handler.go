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

// @BasePath /go-bank-transfer
// @Summary Create Account
// @Description Create new Account in case 'cpf' doesn't exists yet
// @Tags Account
// @Router /accounts [post]
// @Schemes
// @Param AccountRequest body domain.AccountCreatorRequest true "Account request"
// @Accept json
// @Produce json
// @Success 200 {object} domain.AccountCreatorResponse
func (handler *accountHandler) createAccount(context *gin.Context) {
	var body domain.AccountCreatorRequest
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusPreconditionFailed, domain.BuildResponseFromError(err))
		return
	}

	response, err := handler.accountUs.CreateAccount(body)
	if err != nil {
		context.AbortWithStatusJSON(
			domain.GetErrorStatusCode(err),
			domain.BuildResponseFromError(err),
		)
		return
	}

	context.SecureJSON(http.StatusOK, response)
}
