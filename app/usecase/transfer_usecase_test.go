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
	"time"

	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/app/usecase/mocks"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TransferBetweenAccounts

func TestTransferBetweenAccounts_Success(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	originUuid := uuid.New()
	destinationUuid := uuid.New()
	transferId := uuid.New()
	amount := 159.47

	transfer := domain.TransferCreatorResponse{
		Id:                           transferId.String(),
		OldAccountOriginBalance:      159.47,
		NewAccountOriginBalance:      0,
		OldAccountDestinationBalance: 0,
		NewAccountDestinationBalance: 159.47,
	}

	mockRepo.On("TransferBetweenAccounts", mock.Anything, mock.Anything, mock.Anything).
		Return(&transfer, nil)

	response, err := usecase.TransferBetweenAccounts(originUuid.String(), domain.TransferRequest{
		AccountDestinationId: destinationUuid.String(),
		Amount:               amount,
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, transferId.String(), response.Id)
	assert.Equal(t, amount, response.NewAccountDestinationBalance)
	mockRepo.AssertExpectations(t)
}

func TestTransferBetweenAccounts_InvalidAccountOriginId(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	expectedErr := domain.ErrInvalidAccountId

	originUuid := "uuid.New()"
	destinationUuid := uuid.New()
	amount := 159.47

	response, err := usecase.TransferBetweenAccounts(originUuid, domain.TransferRequest{
		AccountDestinationId: destinationUuid.String(),
		Amount:               amount,
	})

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestTransferBetweenAccounts_InvalidAccountDestinationId(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	expectedErr := domain.ErrInvalidAccountId

	originUuid := uuid.New()
	destinationUuid := "uuid.New()"
	amount := 159.47

	response, err := usecase.TransferBetweenAccounts(originUuid.String(),
		domain.TransferRequest{
			AccountDestinationId: destinationUuid,
			Amount:               amount,
		},
	)

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestTransferBetweenAccounts_ErrorOnTransferingFunds(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	expectedErr := errors.New("error transfering funds")

	originUuid := uuid.New()
	destinationUuid := uuid.New()
	amount := 159.47

	mockRepo.On("TransferBetweenAccounts", mock.Anything, mock.Anything, mock.Anything).
		Return(&domain.TransferCreatorResponse{}, expectedErr)

	_, err := usecase.TransferBetweenAccounts(originUuid.String(),
		domain.TransferRequest{
			AccountDestinationId: destinationUuid.String(),
			Amount:               amount,
		})

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	mockRepo.AssertExpectations(t)
}

// GetAccountOriginTransfers

func TestGetAccountOriginTransfers_Success(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	accountOriginId := uuid.New()

	transfers := []domain.Transfer{
		{
			Id:                   uuid.NewString(),
			AccountOriginId:      accountOriginId.String(),
			AccountDestinationId: uuid.NewString(),
			Amount:               123.50,
			CreatedAt:            time.Now().String(),
		},
		{
			Id:                   uuid.NewString(),
			AccountOriginId:      accountOriginId.String(),
			AccountDestinationId: uuid.NewString(),
			Amount:               50.10,
			CreatedAt:            time.Now().String(),
		},
	}

	mockRepo.On("GetAccountOriginTransfers", mock.AnythingOfType("uuid.UUID")).
		Return(transfers, nil)

	response, err := usecase.GetAccountOriginTransfers(accountOriginId.String())

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, len(transfers), len(response.Transfers))
	mockRepo.AssertExpectations(t)
}

func TestGetAccountOriginTransfers_InvalidOriginId(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	accountOriginId := "uuid.New()"
	expectedErr := domain.ErrInvalidAccountId

	response, err := usecase.GetAccountOriginTransfers(accountOriginId)

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestGetAccountOriginTransfers_ErrorRetrievingFromDb(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	usecase := NewTransferUsecase(mockRepo)

	accountOriginId := uuid.New()
	expectedErr := domain.ErrInvalidAccountId

	mockRepo.On("GetAccountOriginTransfers", mock.AnythingOfType("uuid.UUID")).
		Return([]domain.Transfer{}, expectedErr)

	response, err := usecase.GetAccountOriginTransfers(accountOriginId.String())

	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}
