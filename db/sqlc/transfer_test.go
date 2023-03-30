package db

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateTransfer_FindTransferById_FindLastTransferInsertedId(t *testing.T) {
	accounts := GetNewRandomAccountParams(2)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	transferParams := CreateTransferParams{
		ID:            uuid.NewV4().String(),
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        50,
	}

	err := testQueries.CreateTransfer(context.Background(), transferParams)
	require.NoError(t, err)

	transferInserted, err := testQueries.FindTransferById(context.Background(), transferParams.ID)
	require.NoError(t, err)

	require.NotZero(t, transferInserted.ID)
	require.NotZero(t, transferInserted.CreatedAt)

	require.NotEmpty(t, transferInserted)
	require.Equal(t, transferParams.Amount, transferInserted.Amount)
	require.Equal(t, transferParams.FromAccountID, transferInserted.FromAccountID)
	require.Equal(t, transferParams.ToAccountID, transferInserted.ToAccountID)
}

func TestQueries_FindAllTransfers(t *testing.T) {
	accounts := GetNewRandomAccountParams(2)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	transferParams1 := CreateTransferParams{
		ID:            uuid.NewV4().String(),
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        50,
	}

	transferParams2 := CreateTransferParams{
		ID:            uuid.NewV4().String(),
		FromAccountID: accounts[1].ID,
		ToAccountID:   accounts[0].ID,
		Amount:        40,
	}

	err := testQueries.CreateTransfer(context.Background(), transferParams1)
	require.NoError(t, err)

	err = testQueries.CreateTransfer(context.Background(), transferParams2)
	require.NoError(t, err)

	accs, err := testQueries.FindAllTransfers(context.Background())
	require.NoError(t, err)

	require.NotZero(t, len(accs) >= 2)
}

func TestQueries_FindTransfersByFromAccountIdAndToAccountId(t *testing.T) {
	accounts := GetNewRandomAccountParams(2)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
	}

	transferParams1 := CreateTransferParams{
		ID:            uuid.NewV4().String(),
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        50,
	}

	transferParams2 := CreateTransferParams{
		ID:            uuid.NewV4().String(),
		FromAccountID: accounts[1].ID,
		ToAccountID:   accounts[0].ID,
		Amount:        40,
	}

	err := testQueries.CreateTransfer(context.Background(), transferParams1)
	require.NoError(t, err)

	err = testQueries.CreateTransfer(context.Background(), transferParams2)
	require.NoError(t, err)

	args := FindTransfersByFromAccountIdAndToAccountIdParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Limit:         5,
		Offset:        0,
	}

	accs, err := testQueries.FindTransfersByFromAccountIdAndToAccountId(context.Background(), args)
	require.NoError(t, err)

	require.NotZero(t, accs)
}
