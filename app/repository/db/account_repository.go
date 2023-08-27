/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db/queries"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
)

type accountRepository struct {
	dbClient *pgxpool.Pool
}

func NewAccountRepository(dbClient *pgxpool.Pool) domain.AccountRepository {
	return &accountRepository{dbClient: dbClient}
}

func (r *accountRepository) CreateAccount(name, cpf, hashedPassword string) (*domain.AccountCreatorResponse, error) {
	ctx := context.Background()
	defer ctx.Done()

	tx, err := r.dbClient.Begin(ctx)
	if err != nil {
		logrus.Error("error creating transaction - ", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := queries.New(r.dbClient).WithTx(tx)

	_, err = qtx.CreateLogin(ctx, queries.CreateLoginParams{
		Cpf:    cpf,
		Secret: hashedPassword,
	})
	if err != nil {
		logrus.Error("error creating login - ", err)
		return nil, err
	}

	id, err := qtx.CreateAccount(ctx, queries.CreateAccountParams{
		ID:   uuid.New(),
		Name: name,
		Cpf:  cpf,
	})
	if err != nil {
		logrus.Error("error creating account - ", err)
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		logrus.Error("error commitig transaction - ", err)
		return nil, err
	}

	return &domain.AccountCreatorResponse{
		Id: id.String(),
	}, nil
}

func (r *accountRepository) CountAccountByCpf(cpf string) (int64, error) {
	ctx := context.Background()
	defer ctx.Done()

	queries := queries.New(r.dbClient)

	count, err := queries.CountAccountByCpf(ctx, cpf)
	if err != nil {
		logrus.Error("error couting account - ", err)
		return 0, err
	}

	return count, nil
}

func (r *accountRepository) GetAccounts() ([]domain.Account, error) {
	ctx := context.Background()
	defer ctx.Done()

	queries := queries.New(r.dbClient)

	response, err := queries.GetAccounts(ctx)
	if err != nil {
		logrus.Error("error retrieving accounts - ", err)
		return nil, err
	}

	accounts := make([]domain.Account, 0)
	for _, i := range response {
		accounts = append(accounts, domain.Account{
			Id:        i.ID.String(),
			Name:      i.Name,
			Cpf:       i.Cpf,
			Secret:    i.Secret,
			Balance:   microMoneytoMoney(i.Balance),
			CreatedAt: formatDate(i.CreatedAt),
		})
	}

	return accounts, nil
}

func (r *accountRepository) GetAccountBalance(accountId uuid.UUID) (float64, error) {
	ctx := context.Background()
	defer ctx.Done()

	queries := queries.New(r.dbClient)

	accountBalances, err := queries.GetAccountsBalances(ctx, []uuid.UUID{accountId})
	if err != nil {
		logrus.Error("error retrieving account balance - ", err)
		return 0, err
	}

	if len(accountBalances) == 0 {
		logrus.Error("account not found")
		return 0, domain.ErrAccountNotFound
	}

	if len(accountBalances) > 1 {
		logrus.Error("account balance returned more than one account")
		return 0, errors.New("error retrieving account balance")
	}

	return microMoneytoMoney(accountBalances[0].Balance), nil
}
