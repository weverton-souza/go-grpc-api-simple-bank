-- name: CreateTransfer :exec
INSERT INTO transfer (id, from_account_id, to_account_id, amount)
VALUES (?, ?, ?, ?); commit;

-- name: FindLastTransferInsertedId :one
SELECT LAST_INSERT_ID();

-- name: FindTransferById :one
SELECT * FROM transfer e
WHERE e.id = ? LIMIT 1; commit;

-- name: FindAllTransfers :many
SELECT * FROM transfer;

-- name: FindTransfersByFromAccountIdAndToAccountId :many
SELECT * FROM transfer t
WHERE t.from_account_id = ?
   OR t.to_account_id = ?
ORDER BY t.id
LIMIT ? OFFSET ?;
