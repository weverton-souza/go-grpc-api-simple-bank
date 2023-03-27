package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateEntry_FindEntryById_FindLastEntryInsertedId(t *testing.T) {
	accounts := GetNewRandomAccountParams(1)
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	entry := CreateEntryParams{AccountID: ids[0], Amount: 500}

	err := testQueries.CreateEntry(context.Background(), entry)
	require.NoError(t, err)

	lastInsertedId, err := testQueries.FindLastEntryInsertedId(context.Background())
	entryInserted, err := testQueries.FindEntryById(context.Background(), lastInsertedId)
	require.NoError(t, err)

	require.NotZero(t, entryInserted.ID)
	require.NotZero(t, entryInserted.CreatedAt)

	require.NotEmpty(t, entryInserted)
	require.Equal(t, entryInserted.Amount, entryInserted.Amount)
	require.Equal(t, entry.AccountID, entryInserted.AccountID)
}

func TestQueries_FindAllEntries(t *testing.T) {
	accounts := GetNewRandomAccountParams(2)
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	entry1 := CreateEntryParams{AccountID: ids[0], Amount: 500}
	entry2 := CreateEntryParams{AccountID: ids[0], Amount: 500}

	err := testQueries.CreateEntry(context.Background(), entry1)
	require.NoError(t, err)

	err = testQueries.CreateEntry(context.Background(), entry2)
	require.NoError(t, err)

	accs, err := testQueries.FindAllEntries(context.Background())
	require.NoError(t, err)

	require.NotZero(t, len(accs) >= 2)
}

func TestQueries_FindEntriesByAccountId(t *testing.T) {
	accounts := GetNewRandomAccountParams(1)
	ids := make([]int64, 0)

	for _, account := range accounts {
		err := testQueries.CreateAccount(context.Background(), account)
		require.NoError(t, err)
		lastInsertedId, err := testQueries.FindLastAccountInsertedId(context.Background())
		ids = append(ids, lastInsertedId)
	}

	entry1 := CreateEntryParams{AccountID: ids[0], Amount: 150}
	entry2 := CreateEntryParams{AccountID: ids[0], Amount: 897}

	err := testQueries.CreateEntry(context.Background(), entry1)
	require.NoError(t, err)

	err = testQueries.CreateEntry(context.Background(), entry2)
	require.NoError(t, err)

	args := FindEntriesByAccountIdParams{AccountID: ids[0], Limit: 100, Offset: 0}

	accs, err := testQueries.FindEntriesByAccountId(context.Background(), args)
	require.NoError(t, err)

	require.NotZero(t, len(accs) >= 2)
}
