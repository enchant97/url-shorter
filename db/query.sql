-- name: CreateShort :one
INSERT INTO shorts (
    slug, target_url
) VALUES (
    $1, $2
)
RETURNING id;

-- name: GetShortTargetBySlug :one
SELECT target_url FROM shorts
WHERE slug = $1 AND deleted_at IS null;
