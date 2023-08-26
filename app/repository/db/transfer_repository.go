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

type transferRepository struct {
	dbClient *pgxpool.Pool
}

func NewTransferRepository(dbClient *pgxpool.Pool) domain.TransferRepository {
	return &transferRepository{dbClient: dbClient}
}

func (r *transferRepository) TransferBetweenAccounts(
	amount float64, accountOriginId, accountDestinationId uuid.UUID,
) (*domain.TransferCreatorResponse, error) {
	ctx := context.Background()
	defer ctx.Done()

	transferAmount := moneyToMicroMoney(amount)

	tx, err := r.dbClient.Begin(ctx)
	if err != nil {
		logrus.Error("error creating transaction - ", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := queries.New(r.dbClient).WithTx(tx)

	accounts, err := qtx.GetAccountsBalances(ctx, []uuid.UUID{accountOriginId, accountDestinationId})
	if err != nil {
		logrus.Error("error retrieving account balances - ", err)
		return nil, err
	}

	if len(accounts) != 2 {
		logrus.Error("account balances query did not return two accounts")
		return nil, errors.New("account balances query did not return two accounts")
	}

	originBalance, destinationBalance, err := getBalances(accounts, accountOriginId, accountDestinationId)
	if err != nil {
		return nil, err
	}

	if originBalance < transferAmount {
		return nil, domain.ErrInsifficientFunds
	}

	newOriginBalance, newDestinationBalance := transferBalances(
		originBalance, destinationBalance, transferAmount,
	)

	if err := updateAccountBalances(
		ctx, qtx, accountOriginId, accountDestinationId, newOriginBalance, newDestinationBalance,
	); err != nil {
		return nil, err
	}

	response, err := qtx.CreateTransfer(ctx, queries.CreateTransferParams{
		ID:                   uuid.New(),
		AccountOriginID:      accountOriginId,
		AccountDestinationID: accountDestinationId,
		Amount:               transferAmount,
	})
	if err != nil {
		logrus.Error("error creating transfer record - ", err)
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		logrus.Error("error commitig transaction - ", err)
		return nil, err
	}

	return &domain.TransferCreatorResponse{
		Id:                           response.String(),
		OldAccountOriginBalance:      microMoneytoMoney(originBalance),
		NewAccountOriginBalance:      microMoneytoMoney(newOriginBalance),
		OldAccountDestinationBalance: microMoneytoMoney(destinationBalance),
		NewAccountDestinationBalance: microMoneytoMoney(newDestinationBalance),
	}, nil
}

func updateAccountBalances(
	ctx context.Context,
	qtx *queries.Queries,
	accountOriginId, accountDestinationId uuid.UUID,
	newOriginBalance, newDestinationBalance int64,
) error {
	if err := qtx.UpdateAccountBalance(ctx, queries.UpdateAccountBalanceParams{
		ID: accountOriginId, Balance: newOriginBalance,
	}); err != nil {
		logrus.Error("error updating origin account balance - ", err)
		return err
	}

	if err := qtx.UpdateAccountBalance(ctx, queries.UpdateAccountBalanceParams{
		ID: accountDestinationId, Balance: newDestinationBalance,
	}); err != nil {
		logrus.Error("error updating destination account balance - ", err)
		return err
	}

	return nil
}

func (r *transferRepository) GetAccountOriginTransfers(
	accountOriginId uuid.UUID,
) ([]domain.Transfer, error) {
	ctx := context.Background()
	defer ctx.Done()

	queries := queries.New(r.dbClient)

	response, err := queries.GetTransfers(ctx, accountOriginId)
	if err != nil {
		logrus.Error("error retrieving transfers - ", err)
		return nil, err
	}

	transfers := make([]domain.Transfer, 0)
	for _, i := range response {
		transfers = append(transfers, domain.Transfer{
			Id:                   i.ID.String(),
			AccountOriginId:      i.AccountDestinationID.String(),
			AccountDestinationId: i.AccountDestinationID.String(),
			Amount:               microMoneytoMoney(i.Amount),
			CreatedAt:            formatDate(i.CreatedAt),
		})
	}

	return transfers, nil
}
