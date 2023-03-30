package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	accounts := GetNewRandomAccountParams(2)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ID := uuid.NewV4().String()
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				ID:            ID,
				FromAccountID: accounts[0].ID,
				ToAccountID:   accounts[1].ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, accounts[0].ID, transfer.FromAccountID)
		require.Equal(t, accounts[1].ID, transfer.ToAccountID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.FindTransferById(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, accounts[0].ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.FindEntryById(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, accounts[1].ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.FindEntryById(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}
}
