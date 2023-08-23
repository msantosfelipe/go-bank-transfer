/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import "time"

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"string"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   string    `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type AccountCreatorRequest struct {
	Name   string `json:"string"`
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type AccountCreatorResponse struct {
	Id string `json:"id"`
}

// Account usecase methods deifinition
type AccountUsecase interface {
	CreateAccount(accountRequest AccountCreatorRequest) (*AccountCreatorResponse, error)
}

// Account repository methods deifinition
type AccountRepository interface {
	CreateAccount(accountRequest AccountCreatorRequest) (*AccountCreatorResponse, error)
}
