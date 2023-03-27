package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateTransfer_FindTransferById_FindLastTransferInsertedId(t *testing.T) {
	accounts := GetNewRandomAccountParams(2)
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	transferParams := CreateTransferParams{
		FromAccountID: ids[0],
		ToAccountID:   ids[1],
		Amount:        50,
	}

	err := testQueries.CreateTransfer(context.Background(), transferParams)
	require.NoError(t, err)

	lastInsertedId, err := testQueries.FindLastTransferInsertedId(context.Background())
	transferInserted, err := testQueries.FindTransferById(context.Background(), lastInsertedId)
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
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	transferParams1 := CreateTransferParams{
		FromAccountID: ids[0],
		ToAccountID:   ids[1],
		Amount:        50,
	}

	transferParams2 := CreateTransferParams{
		FromAccountID: ids[1],
		ToAccountID:   ids[0],
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
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	transferParams1 := CreateTransferParams{
		FromAccountID: ids[0],
		ToAccountID:   ids[1],
		Amount:        50,
	}

	transferParams2 := CreateTransferParams{
		FromAccountID: ids[1],
		ToAccountID:   ids[0],
		Amount:        40,
	}

	err := testQueries.CreateTransfer(context.Background(), transferParams1)
	require.NoError(t, err)

	err = testQueries.CreateTransfer(context.Background(), transferParams2)
	require.NoError(t, err)

	args := FindTransfersByFromAccountIdAndToAccountIdParams{
		FromAccountID: ids[0],
		ToAccountID:   ids[1],
		Limit:         5,
		Offset:        0,
	}

	accs, err := testQueries.FindTransfersByFromAccountIdAndToAccountId(context.Background(), args)
	require.NoError(t, err)

	require.NotZero(t, accs)
}
