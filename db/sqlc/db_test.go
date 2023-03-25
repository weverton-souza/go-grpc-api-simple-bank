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

	account := CreateAccountParams{
		Owner:    randomNOwner(1)[0],
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}

	qtx := testQueries.WithTx(tx)
	err = qtx.CreateAccount(context.Background(), account)
	if err != nil {
		log.Fatal("Cannot create a new account:", err)
	}

	lastInsertedId, err := qtx.FindLastInsertedId(context.Background())
	accountInserted, err := qtx.FindAccountById(context.Background(), lastInsertedId)

	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, account.Owner, accountInserted.Owner)
	require.Equal(t, account.Balance, accountInserted.Balance)
	require.Equal(t, account.Currency, accountInserted.Currency)
}
