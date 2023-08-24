/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db/queries"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
)

type loginRepository struct {
	dbClient *pgxpool.Pool
}

func NewLoginRepository(dbClient *pgxpool.Pool) domain.LoginRepository {
	return &loginRepository{dbClient: dbClient}
}

func (r *loginRepository) GetLoginByCpf(cpf string) (*domain.Login, error) {
	ctx := context.Background()
	defer ctx.Done()

	queries := queries.New(r.dbClient)

	login, err := queries.GetLogin(ctx, cpf)
	if err != nil {
		logrus.Error("error retrieving login - ", err)
		return nil, err
	}

	return &domain.Login{
		Cpf:    login.Cpf,
		Secret: login.Secret,
	}, nil
}
