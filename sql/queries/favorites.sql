-- name: CreateFavorite :execute
INSERT INTO favorites (user_id, place_id)
VALUES (?, ?);
-- name: DeleteFavorite :execute
DELETE FROM favorites
WHERE user_id = ?
  AND place_id = ?;
-- name: FindFavoritesByUserId :many
SELECT *
FROM favorites
WHERE user_id = ?;