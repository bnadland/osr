-- name: GetItems :many
SELECT * FROM items LIMIT 20;

-- name: CreateFeed :one
INSERT INTO feeds (title, link, updated_at)
VALUES ($1, $2, NOW())
ON CONFLICT (link)
DO UPDATE SET updated_at=NOW()
RETURNING *;

-- name: CreateItem :one
INSERT INTO items (feed_id, title, link, updated_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (link)
DO UPDATE SET updated_at=NOW()
RETURNING *;