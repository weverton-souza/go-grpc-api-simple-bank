// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :exec
INSERT INTO entry (id, account_id, amount)
VALUES (?, ?, ?)
`

type CreateEntryParams struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) error {
	_, err := q.db.ExecContext(ctx, createEntry, arg.ID, arg.AccountID, arg.Amount)
	return err
}

const findAllEntries = `-- name: FindAllEntries :many
SELECT id, account_id, amount, created_at FROM entry
`

func (q *Queries) FindAllEntries(ctx context.Context) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, findAllEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findEntriesByAccountId = `-- name: FindEntriesByAccountId :many
SELECT id, account_id, amount, created_at FROM entry e
WHERE e.account_id = ?
ORDER BY e.id
LIMIT ? OFFSET ?
`

type FindEntriesByAccountIdParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) FindEntriesByAccountId(ctx context.Context, arg FindEntriesByAccountIdParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, findEntriesByAccountId, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findEntryById = `-- name: FindEntryById :one
SELECT id, account_id, amount, created_at FROM entry e
WHERE e.id = ? LIMIT 1
`

func (q *Queries) FindEntryById(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, findEntryById, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const findLastEntryInsertedId = `-- name: FindLastEntryInsertedId :one
SELECT LAST_INSERT_ID()
`

func (q *Queries) FindLastEntryInsertedId(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, findLastEntryInsertedId)
	var last_insert_id int64
	err := row.Scan(&last_insert_id)
	return last_insert_id, err
}
