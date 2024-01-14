-- name: GetItems :many
SELECT i.item_id, i.title, i.link, i.published_at, f.title AS feed_title
FROM items i
LEFT JOIN feeds f ON f.feed_id=i.feed_id
ORDER BY i.published_at DESC
LIMIT 20;

-- name: CreateFeed :one
INSERT INTO feeds (title, link, updated_at)
VALUES ($1, $2, NOW())
ON CONFLICT (link)
DO UPDATE SET updated_at=NOW()
RETURNING *;

-- name: CreateItem :one
INSERT INTO items (feed_id, title, link, categories, content, published_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW())
ON CONFLICT (link)
DO UPDATE SET updated_at=NOW()
RETURNING *;