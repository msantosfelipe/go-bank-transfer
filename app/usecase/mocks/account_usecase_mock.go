/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package mocks

import (
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/stretchr/testify/mock"
)

// Mock of account repository for testing
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) CountAccountByCpf(cpf string) (int64, error) {
	args := m.Called(cpf)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockAccountRepository) CreateAccount(name, cpf, password string) (*domain.AccountCreatorResponse, error) {
	args := m.Called(name, cpf, password)
	return args.Get(0).(*domain.AccountCreatorResponse), args.Error(1)
}

func (m *MockAccountRepository) GetAccounts() ([]domain.Account, error) {
	args := m.Called()
	return args.Get(0).([]domain.Account), args.Error(1)
}
