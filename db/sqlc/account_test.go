package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestQueries_CreateAccount_and_FindLastInsertedId_and_FindAccountById(t *testing.T) {
	accounts := GetNewRandomAccountParams(1)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	accountInserted, err := testQueries.FindAccountById(context.Background(), accounts[0].ID)

	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, accounts[0].Owner, accountInserted.Owner)
	require.Equal(t, accounts[0].Balance, accountInserted.Balance)
	require.Equal(t, accounts[0].Currency, accountInserted.Currency)
}

func TestQueries_UpdateAccount(t *testing.T) {
	accounts := GetNewRandomAccountParams(1)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	accountInserted, err := testQueries.FindAccountById(context.Background(), accounts[0].ID)
	require.NoError(t, err)

	updateAccountParams := UpdateAccountParams{ID: accounts[0].ID, Balance: int64(rand.Intn(10000-0) + 0)}

	err = testQueries.UpdateAccount(context.Background(), updateAccountParams)
	require.NoError(t, err)

	accountUpdated, err := testQueries.FindAccountById(context.Background(), accounts[0].ID)
	require.NoError(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)

	require.NotEmpty(t, accountInserted)
	require.Equal(t, accounts[0].Owner, accountInserted.Owner)
	require.NotEqual(t, accounts[0].Balance, accountUpdated.Balance)
	require.Equal(t, accounts[0].Balance, accountInserted.Balance)
	require.Equal(t, accounts[0].Currency, accountInserted.Currency)
}

func TestQueries_FindAllAccounts(t *testing.T) {
	accounts := GetNewRandomAccountParams(3)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	accs, err := testQueries.FindAllAccounts(context.Background())
	require.NoError(t, err)

	require.NotZero(t, len(accs))
}

func TestQueries_DeleteAccount(t *testing.T) {
	accounts := GetNewRandomAccountParams(1)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	accountInserted, err := testQueries.FindAccountById(context.Background(), accounts[0].ID)
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), accounts[0].ID)
	require.NoError(t, err)

	accountDeleted, err := testQueries.FindAccountById(context.Background(), accounts[0].ID)
	require.Error(t, err)

	require.NotZero(t, accountInserted.ID)
	require.NotZero(t, accountInserted.CreatedAt)
	require.NotEmpty(t, accountInserted)

	require.Zero(t, accountDeleted.ID)
	require.Zero(t, accountDeleted.CreatedAt)
	require.Empty(t, accountDeleted)
}
