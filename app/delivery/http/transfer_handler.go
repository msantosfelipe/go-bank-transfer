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
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/jwt"
)

type transferHandler struct {
	transferUs domain.TransferUsecase
}

// NewLoginRouter handle REST requests
func NewTransferHandler(router *gin.RouterGroup, transferUs domain.TransferUsecase) {
	handler := transferHandler{
		transferUs: transferUs,
	}

	router.POST("/transfers", handler.createTransfer)
}

// @BasePath /go-bank-transfer
// @Summary Transfer amount
// @Description Transfer amount from origin account to destination account
// @Tags Transfers
// @Router /transfers [post]
// @Param TransferRequest body domain.TransferRequest true "Transfer request"
// @Param Authorization header string true "Token"
// @Accept json
// @Produce json
// @Success 201 {object} domain.TransferCreatorResponse
// @Failure 400 {object} domain.ResponseError
// @Failure 500 {object} domain.ResponseError
func (handler *transferHandler) createTransfer(context *gin.Context) {
	originAccountId, err := jwt.ExtractAccountOriginId(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, domain.BuildResponseFromError(err))
		return
	}

	var body domain.TransferRequest
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, domain.BuildResponseFromError(err))
		return
	}

	response, err := handler.transferUs.TransferBetweenAccounts(originAccountId, body)
	if err != nil {
		context.AbortWithStatusJSON(
			domain.GetErrorStatusCode(err),
			domain.BuildResponseFromError(err),
		)
		return
	}

	context.SecureJSON(http.StatusCreated, response)
}
