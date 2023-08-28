/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/msantosfelipe/go-bank-transfer/config"
	"github.com/sirupsen/logrus"
)

const retries = 5

func InitDb(ctx context.Context) *pgxpool.Pool {
	for i := 1; i <= retries; i++ {
		dbClient, err := newDbClient(ctx)
		if err != nil {
			time.Sleep(1 * time.Second)
			logrus.Warn(fmt.Sprintf("error connecting to database retry %v of %v", i, retries))
			continue
		}

		if err := dbClient.Ping(ctx); err != nil {
			logrus.Error("error testing db connection - ", err)
			break
		}

		return dbClient
	}

	panic("failed to connect do database")
}

func newDbClient(ctx context.Context) (*pgxpool.Pool, error) {
	devDbClient, err := initDevDb(ctx)
	if err != nil {
		return initLocalDb(ctx)
	}

	return devDbClient, nil
}

func initDevDb(ctx context.Context) (*pgxpool.Pool, error) {
	dbClient, err := pgxpool.Connect(ctx, getDbUri(config.ENV.DbHost))
	if err != nil {
		logrus.Warn("error connecting to dev db - ", err)
		return nil, err
	}

	logrus.Info("successfully connected to dev db")
	return dbClient, nil
}

func initLocalDb(ctx context.Context) (*pgxpool.Pool, error) {
	dbClient, err := pgxpool.Connect(ctx, getDbUri(config.ENV.DbHostLocal))
	if err != nil {
		logrus.Warn("error connecting to localhost db - ", err)
		return nil, err
	}

	logrus.Info("successfully connected to localhost db")
	return dbClient, nil
}

func getDbUri(hostname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.ENV.DbUser,
		config.ENV.DbPass,
		hostname,
		config.ENV.DbPort,
		config.ENV.DbName,
	)
}
