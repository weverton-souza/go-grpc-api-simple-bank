package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestQueries_WithTx(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	tx, err := conn.Begin()
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Fatal("Cannot connect to db:", err)
		}
	}(tx)

	accounts := GetNewRandomAccountParams(1)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	qtx := testQueries.WithTx(tx)
	err = qtx.CreateAccount(context.Background(), accounts[0])
	if err != nil {
		log.Fatal("Cannot create a new account:", err)
	}

	lastInsertedId, err := qtx.FindLastAccountInsertedId(context.Background())
	accountInserted, err := qtx.FindAccountById(context.Background(), lastInsertedId)

	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, accounts[0].Owner, accountInserted.Owner)
	require.Equal(t, accounts[0].Balance, accountInserted.Balance)
	require.Equal(t, accounts[0].Currency, accountInserted.Currency)
}
