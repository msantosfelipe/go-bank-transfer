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
	// ErrAccountConflict will throw if the cpf is invalid
	ErrAccountInvalidCpf = errors.New("invalid cpf")
	// ErrAccountConflict will throw if the current account already exists
	ErrAccountConflict = errors.New("already exists an account for this cpf")
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
	case ErrAccountInvalidCpf:
		return http.StatusPreconditionFailed
	case ErrAccountConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
