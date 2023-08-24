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
	router.GET("/accounts", handler.GetAccounts)
	router.GET("/accounts/:account_id/balance", handler.GetAccountBalance)
}

// @BasePath /go-bank-transfer
// @Summary Create account
// @Description Create new Account in case 'cpf' doesn't exists yet
// @Tags Account
// @Router /accounts [post]
// @Schemes
// @Param AccountRequest body domain.AccountCreatorRequest true "Account request"
// @Accept json
// @Produce json
// @Success 201 {object} domain.AccountCreatorResponse
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

	context.SecureJSON(http.StatusCreated, response)
}

// @BasePath /go-bank-transfer
// @Summary Get accounts
// @Description Return the list of all accounts. Fields 'secret' and 'balance' are omitted
// @Tags Account
// @Router /accounts [get]
// @Schemes
// @Produce json
// @Success 200 {object} domain.AccountList
func (handler *accountHandler) GetAccounts(context *gin.Context) {
	response, err := handler.accountUs.GetAccounts()
	if err != nil {
		context.AbortWithStatusJSON(
			domain.GetErrorStatusCode(err),
			domain.BuildResponseFromError(err),
		)
		return
	}

	context.SecureJSON(http.StatusOK, response)
}

// @BasePath /go-bank-transfer
// @Summary Get account balance
// @Description Return account balance
// @Tags Account
// @Router /accounts/{account_id}/balance [get]
// @Param account_id path string true "Account ID"
// @Schemes
// @Produce json
// @Success 200 {object} domain.AccountBalance
func (handler *accountHandler) GetAccountBalance(context *gin.Context) {
	accountId, exists := context.Params.Get("account_id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusPreconditionFailed,
			domain.BuildResponseFromError(domain.ErrBadParamInput))
		return
	}

	response, err := handler.accountUs.GetAccountBalance(accountId)
	if err != nil {
		context.AbortWithStatusJSON(
			domain.GetErrorStatusCode(err),
			domain.BuildResponseFromError(err),
		)
		return
	}

	context.SecureJSON(http.StatusOK, response)
}
