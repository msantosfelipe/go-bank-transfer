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

// Login usecase methods deifinition
type LoginUsecase interface {
	AuthenticateUser(login Login) error
}

// Login repository methods deifinition
type LoginRepository interface {
	GetLoginByCpf(cpf string) (*Login, error)
}
