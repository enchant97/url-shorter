-- name: CreateShort :one
INSERT INTO shorts (
    slug, target_url
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetShortTargetBySlug :one
SELECT target_url FROM shorts
WHERE slug = $1 AND deleted_at IS null;

-- name: GetShortByID :one
SELECT * FROM shorts
WHERE id = $1 AND deleted_at IS null;

-- name: LatestShorts :many
SELECT * FROM shorts
WHERE deleted_at IS null
LIMIT $1;

-- name: UpdateShortByID :one
UPDATE shorts
SET updated_at = default,
    target_url = $2
WHERE id = $1 AND deleted_at IS null
RETURNING *;
