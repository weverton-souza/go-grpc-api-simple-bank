-- name: CreateAccount :execresult
INSERT INTO accounts (id, owner, balance, currency) VALUES (?, ?, ?, ?);
