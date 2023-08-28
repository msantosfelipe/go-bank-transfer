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
	// ErrInvalidPasswordLength will throw if the user pass a too short or too long password
	ErrInvalidPasswordLength = errors.New("the password must have between 6 and 16 characteres")
	// ErrAccountConflict will throw if the search query returns will no results
	ErrNoRowsInResultSet = errors.New("no rows in result set")
	// ErrUserNotAuthorized will throw if the uses tries to access a protected resource with an invalid token
	ErrUserNotAuthorized = errors.New("user not authorized")
	// ErrInsifficientFunds will throw if the origin account has insifficient funds to make the transfer
	ErrInsifficientFunds = errors.New("origin account does not have sufficient funds for transfer")
	// ErrInsifficientFunds will throw if the user tries to pass an invalid account id
	ErrInvalidAccountId = errors.New("invalid account id")
	// ErrInsifficientFunds will throw if the account was not found in database
	ErrAccountNotFound = errors.New("account not found")
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
	case ErrInsifficientFunds:
		return http.StatusBadRequest
	case ErrInvalidAccountId:
		return http.StatusBadRequest
	case ErrAccountNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
