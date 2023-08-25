/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/jwt"
	"github.com/sirupsen/logrus"
)

// LogMiddleware base logger for handler requests
func LogMiddleware(context *gin.Context) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"path":         context.Request.URL.Path,
		"x-request-id": uuid.New().String(),
		"date_time":    time.Now(),
	})

	logger.Info("Request initialized.")
	context.Set("log", logger)
	context.Next()
}

// JwthMiddleware handle jwt token validation
func JwthMiddleware(context *gin.Context) {
	if ok := jwt.IsTokenValid(context); !ok {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			domain.BuildResponseFromError(domain.ErrUserNotAuthorized),
		)
		context.Abort()
		return
	}
	context.Next()
}
