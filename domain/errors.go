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
	// ErrAccountConflict will throw if the param in url is not a valid param
	ErrBadParamInput = errors.New("given param is not valid")
	// ErrAccountConflict will throw if the user tries to log in with an invalid login or password
	ErrInvalidLogin = errors.New("login or password is not valid")
	// ErrAccountConflict will throw if the search query returns will no results
	ErrNoRowsInResultSet = errors.New("no rows in result set")
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
	case ErrInvalidLogin:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
