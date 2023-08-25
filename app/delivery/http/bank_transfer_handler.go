/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/jwt"
)

type bankTransferHandler struct {
}

// NewLoginRouter handle REST requests
func NewBankTransferHandler(router *gin.RouterGroup) {
	handler := bankTransferHandler{}

	router.POST("/transfers", handler.createTransfer)
}

func (handler *bankTransferHandler) createTransfer(context *gin.Context) {
	id, err := jwt.ExtractAccountOriginId(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, domain.BuildResponseFromError(err))
		return
	}
	fmt.Println(id)

	// var body domain.Login
	// if err := context.BindJSON(&body); err != nil {
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, domain.BuildResponseFromError(err))
	// 	return
	// }

	context.SecureJSON(http.StatusOK, "ok")
}
