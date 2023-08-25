/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package mocks

import (
	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/stretchr/testify/mock"
)

// Mock of transfer repository for testing
type MockTransferRepository struct {
	mock.Mock
}

func (m *MockTransferRepository) TransferBetweenAccounts(
	amount float64, accountOriginId, accountDestinationId uuid.UUID,
) (*domain.TransferCreatorResponse, error) {
	args := m.Called(amount, accountOriginId, accountDestinationId)
	return args.Get(0).(*domain.TransferCreatorResponse), args.Error(1)
}
