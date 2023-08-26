/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import (
	"github.com/google/uuid"
)

const DefaultBalance = 0

// Account content struct deifinition
type Account struct {
	Id        string  `json:"id"`
	Name      string  `json:"string"`
	Cpf       string  `json:"cpf"`
	Secret    string  `json:"-"`
	Balance   float64 `json:"-"`
	CreatedAt string  `json:"created_at"`
}

type AccountList struct {
	Accounts []Account `json:"accounts"`
}

type AccountCreatorRequest struct {
	Name   string `json:"name"`
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type AccountCreatorResponse struct {
	Id string `json:"id"`
}

type AccountBalance struct {
	Balance float64 `json:"balance"`
}

// Account usecase methods deifinition
type AccountUsecase interface {
	CreateAccount(request AccountCreatorRequest) (*AccountCreatorResponse, error)
	GetAccounts() (*AccountList, error)
	GetAccountBalance(accountId string) (*AccountBalance, error)
}

// Account repository methods deifinition
type AccountRepository interface {
	CreateAccount(name, cpf, hashedPassword string) (*AccountCreatorResponse, error)
	CountAccountByCpf(cpf string) (int64, error)
	GetAccounts() ([]Account, error)
	GetAccountBalance(accountId uuid.UUID) (float64, error)
}
