/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

// Login content struct deifinition
type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type JwtToken struct {
	Token string `json:"token"`
}

// Login usecase methods deifinition
type LoginUsecase interface {
	AuthenticateUser(credentials Login) (*JwtToken, error)
}

// Login repository methods deifinition
type LoginRepository interface {
	GetLoginAndAccount(cpf string) (*Login, string, error)
}
