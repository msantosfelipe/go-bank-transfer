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

// Mock of login repository for testing
type MockLoginRepository struct {
	mock.Mock
}

func (m *MockLoginRepository) GetLoginAndAccount(cpf string) (*domain.Login, string, error) {
	args := m.Called(cpf)
	return args.Get(0).(*domain.Login), args.String(1), args.Error(2)
}
