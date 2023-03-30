-- name: CreateAccount :exec
INSERT INTO account (id, owner, balance, currency)
VALUES (?, ?, ?, ?); commit;

-- name: FindAccountById :one
SELECT * FROM account acc
WHERE acc.id = ? LIMIT 1; commit;

-- name: FindAllAccounts :many
SELECT * FROM account;

-- name: UpdateAccount :exec
UPDATE account SET balance = ? WHERE id = ?; commit;

-- name: DeleteAccount :exec
DELETE FROM account WHERE id = ?; commit;
