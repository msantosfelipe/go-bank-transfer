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

	"github.com/google/uuid"
)

var (
	// ErrAccountConflict will throw if the cpf is invalid
	ErrAccountInvalidCpf = errors.New("invalid cpf")
	// ErrAccountConflict will throw if the current account already exists
	ErrAccountConflict = errors.New("already exists an account for this cpf")
	ErrBadParamInput   = errors.New("given param is not valid")
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

	if uuid.IsInvalidLengthError(err) {
		return http.StatusBadRequest
	}

	switch err {
	case ErrAccountInvalidCpf:
		return http.StatusBadRequest
	case ErrAccountConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
