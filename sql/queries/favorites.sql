-- name: CreateFavorite :exec
INSERT INTO favorites (user_id, place_id)
VALUES (?, ?);
-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id = ?
  AND place_id = ?;
-- name: FindFavoritesByUserId :many
SELECT *
FROM favorites
WHERE user_id = ?;