/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import "time"

const DefaultBalance = 0

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"string"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"-"`
	Balance   float64   `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
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

// Account usecase methods deifinition
type AccountUsecase interface {
	CreateAccount(request AccountCreatorRequest) (*AccountCreatorResponse, error)
	GetAccounts() (*AccountList, error)
}

// Account repository methods deifinition
type AccountRepository interface {
	CreateAccount(name, cpf, hashedPassword string) (*AccountCreatorResponse, error)
	CountAccountByCpf(cpf string) (int64, error)
	GetAccounts() ([]Account, error)
}
