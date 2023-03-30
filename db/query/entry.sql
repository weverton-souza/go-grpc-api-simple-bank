-- name: CreateEntry :exec
INSERT INTO entry (id, account_id, amount)
VALUES (?, ?, ?); commit;

-- name: FindEntryById :one
SELECT * FROM entry e
WHERE e.id = ? LIMIT 1; commit;

-- name: FindAllEntries :many
SELECT * FROM entry;

-- name: FindEntriesByAccountId :many
SELECT * FROM entry e
WHERE e.account_id = ?
ORDER BY e.id
LIMIT ? OFFSET ?;
