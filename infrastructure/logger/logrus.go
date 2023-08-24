/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// LogMiddleware base logger
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
