/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase/mocks"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccount_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)
	id := uuid.New()

	mockRepo.On("CountAccountByCpf", mock.Anything).Return(0, nil)
	mockRepo.On("CreateAccount", mock.Anything, mock.Anything, mock.Anything).Return(&domain.AccountCreatorResponse{
		Id: id.String(),
	}, nil)

	request := domain.AccountCreatorRequest{
		Name:   "James Bond",
		Cpf:    "12345678901",
		Secret: "password",
	}

	// Action
	response, err := usecase.CreateAccount(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, id.String(), response.Id)

	mockRepo.AssertExpectations(t)
}

func TestCreateAccount_InvalidCpfLenght(t *testing.T) {
	mockRepo := new(mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	request := domain.AccountCreatorRequest{
		Name:   "James Bond",
		Cpf:    "1234567890",
		Secret: "password",
	}

	response, err := usecase.CreateAccount(request)
	assert.ErrorIs(t, err, domain.ErrAccountInvalidCpf)
	assert.Nil(t, response)

	mockRepo.AssertExpectations(t)
}

func TestCreateAccount_InvalidCpfDigits(t *testing.T) {
	mockRepo := new(mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	request := domain.AccountCreatorRequest{
		Name:   "James Bond",
		Cpf:    "123.456.789-01",
		Secret: "password",
	}

	response, err := usecase.CreateAccount(request)
	assert.ErrorIs(t, err, domain.ErrAccountInvalidCpf)
	assert.Nil(t, response)

	mockRepo.AssertExpectations(t)
}

func TestCreateAccount_CpfAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	mockRepo.On("CountAccountByCpf", mock.Anything).Return(1, nil)

	request := domain.AccountCreatorRequest{
		Name:   "James Bond",
		Cpf:    "12345678901",
		Secret: "password",
	}

	response, err := usecase.CreateAccount(request)
	assert.ErrorIs(t, err, domain.ErrAccountConflict)
	assert.Nil(t, response)

	mockRepo.AssertExpectations(t)
}

func TestCreateAccount_CountByCpfError(t *testing.T) {
	mockRepo := new(mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	expectedErr := errors.New("error in db")

	mockRepo.On("CountAccountByCpf", mock.Anything).Return(0, expectedErr)

	request := domain.AccountCreatorRequest{
		Name:   "James Bond",
		Cpf:    "12345678901",
		Secret: "password",
	}

	response, err := usecase.CreateAccount(request)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, response)

	mockRepo.AssertExpectations(t)
}