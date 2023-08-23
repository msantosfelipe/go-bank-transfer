/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import (
	"errors"
	"net/http"
)

var (
	// ErrConflict will throw if the current record already exists
	ErrConflict = errors.New("your item already exist")
)

// ResponseError struct definition
type ResponseError struct {
	Message string `json:"message"`
}

// BuildResponseFromError --
func BuildResponseFromError(err error) ResponseError {
	msgError := err.Error()
	return ResponseError{
		Message: msgError,
	}
}

// GetErrorStatusCode --
func GetErrorStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
