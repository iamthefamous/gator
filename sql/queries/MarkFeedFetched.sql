-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = now(),
    updated_at = now()
where
    id = $1
RETURNING *;
