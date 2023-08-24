/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package db

import "github.com/msantosfelipe/go-bank-transfer/domain"

type loginRepository struct {
}

func NewLoginRepository() domain.LoginRepository {
	return &loginRepository{}
}

func (r *loginRepository) GetLoginByCpf(cpf string) (*domain.Login, error) {
	return nil, nil
}
