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
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticateUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockLoginRepository)
	usecase := NewLoginUsecase(mockRepo)

	accountId := uuid.New().String()
	cpf := "87832842067"
	secret := "secret"

	credentials := domain.Login{
		Cpf:    cpf,
		Secret: secret,
	}

	hashedSecret, err := crypto.HashPassword(credentials.Secret)
	if err != nil {
		panic(err)
	}

	storedLogin := domain.Login{
		Cpf:    cpf,
		Secret: hashedSecret,
	}

	mockRepo.On("GetLoginAndAccount", mock.Anything).Return(&storedLogin, accountId, nil)

	jwtToken, err := usecase.AuthenticateUser(credentials)

	assert.NoError(t, err)
	assert.NotNil(t, jwtToken)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_LoginDoesNotExists(t *testing.T) {
	mockRepo := new(mocks.MockLoginRepository)
	usecase := NewLoginUsecase(mockRepo)

	expectedError := domain.ErrInvalidLogin
	credentials := domain.Login{
		Cpf:    "87832842067",
		Secret: "secret",
	}

	mockRepo.On("GetLoginAndAccount", mock.Anything).Return(&domain.Login{}, "", domain.ErrNoRowsInResultSet)

	jwtToken, err := usecase.AuthenticateUser(credentials)
	assert.ErrorIs(t, err, expectedError)
	assert.Nil(t, jwtToken)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_ErrorRetrievingLogin(t *testing.T) {
	mockRepo := new(mocks.MockLoginRepository)
	usecase := NewLoginUsecase(mockRepo)

	expectedError := errors.New("error in db")
	credentials := domain.Login{
		Cpf:    "87832842067",
		Secret: "secret",
	}

	mockRepo.On("GetLoginAndAccount", mock.Anything).Return(&domain.Login{}, "", expectedError)

	jwtToken, err := usecase.AuthenticateUser(credentials)
	assert.ErrorIs(t, err, expectedError)
	assert.Nil(t, jwtToken)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_PasswordsDoNotMatch(t *testing.T) {
	mockRepo := new(mocks.MockLoginRepository)
	usecase := NewLoginUsecase(mockRepo)

	accountId := uuid.New().String()
	cpf := "87832842067"
	secret := "secret"

	credentials := domain.Login{
		Cpf:    cpf,
		Secret: secret,
	}

	storedSecret := "secret2"
	hashedSecret, err := crypto.HashPassword(storedSecret)
	if err != nil {
		panic(err)
	}

	storedLogin := domain.Login{
		Cpf:    cpf,
		Secret: hashedSecret,
	}

	mockRepo.On("GetLoginAndAccount", mock.Anything).Return(&storedLogin, accountId, nil)

	jwtToken, err := usecase.AuthenticateUser(credentials)
	assert.ErrorIs(t, err, domain.ErrInvalidLogin)
	assert.Nil(t, jwtToken)
	mockRepo.AssertExpectations(t)
}
