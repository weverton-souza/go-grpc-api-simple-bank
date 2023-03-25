-- name: CreateAccount :exec
INSERT INTO accounts (id, owner, balance, currency)
VALUES (?, ?, ?, ?); commit;

-- name: FindLastInsertedId :one
SELECT LAST_INSERT_ID();

-- name: FindAccountById :one
SELECT * FROM accounts acc
WHERE acc.id = ? LIMIT 1; commit;

-- name: FindAllAccounts :many
SELECT * FROM accounts;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = ? WHERE id = ?; commit;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = ?; commit;