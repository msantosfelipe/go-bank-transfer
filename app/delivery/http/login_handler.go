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

type loginHandler struct {
	loginUs domain.LoginUsecase
}

// NewLoginRouter handle REST requests
func NewLoginRouter(router *gin.RouterGroup, loginUs domain.LoginUsecase) {
	handler := loginHandler{
		loginUs: loginUs,
	}

	router.Use(
		logger.LogMiddleware,
	)

	router.POST("/login", handler.authenticate)
}

// @BasePath /go-bank-transfer
// @Summary Login
// @Description Authemticate user
// @Tags Login
// @Router /login [post]
// @Schemes
// @Param Login body domain.Login true "Login request"
// @Accept json
// @Produce json
// @Success 201 {object} domain.JwtToken
func (handler *loginHandler) authenticate(context *gin.Context) {
	var body domain.Login
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, domain.BuildResponseFromError(err))
		return
	}

	response, err := handler.loginUs.AuthenticateUser(body)
	if err != nil {
		context.AbortWithStatusJSON(
			domain.GetErrorStatusCode(err),
			domain.BuildResponseFromError(err),
		)
		return
	}

	context.SecureJSON(http.StatusOK, response)
}
