package db

import (
	"context"
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Store struct {
	*Queries
	db *sql.DB
}

type TransferTxParams struct {
	ID            string `json:"id"`
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

type TransferTxResult struct {
	Transfer      Transfer `json:"transfer"`
	FromAccountID string   `json:"from_account_id"`
	ToAccountID   string   `json:"to_account_id"`
	Amount        int64    `json:"amount"`
	FromEntry     Entry    `json:"from_entry"`
	ToEntry       Entry    `json:"to_entry"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, Rollback Error: %v", err, rbErr)
		}
	}

	return tx.Commit()
}

func (s *Store) TransferTx(ctx context.Context, params TransferTxParams) (result TransferTxResult, err error) {

	err = s.execTx(ctx, func(q *Queries) error {

		transferTxParams := CreateTransferParams{
			ID:            uuid.NewV4().String(),
			FromAccountID: params.FromAccountID,
			ToAccountID:   params.ToAccountID,
			Amount:        params.Amount,
		}

		err = q.CreateTransfer(ctx, transferTxParams)
		if err != nil {
			return err
		}

		result.Transfer, err = q.FindTransferById(ctx, transferTxParams.ID)
		if err != nil {
			return err
		}

		fromEntryParams := CreateEntryParams{
			ID:        uuid.NewV4().String(),
			AccountID: params.FromAccountID,
			Amount:    -params.Amount,
		}

		err = q.CreateEntry(ctx, fromEntryParams)
		if err != nil {
			return err
		}

		result.FromEntry, err = q.FindEntryById(ctx, fromEntryParams.ID)
		if err != nil {
			return err
		}

		toEntryParams := CreateEntryParams{
			ID:        uuid.NewV4().String(),
			AccountID: params.ToAccountID,
			Amount:    params.Amount,
		}

		err = q.CreateEntry(ctx, toEntryParams)
		if err != nil {
			return err
		}

		result.ToEntry, err = q.FindEntryById(ctx, toEntryParams.ID)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
