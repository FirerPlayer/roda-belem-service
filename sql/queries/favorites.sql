-- name: CreateFavorite :exec
INSERT INTO favorites (user_id, place_id)
VALUES (?, ?);
-- name: DeleteFavoriteByUserIdAndPlaceId :exec
DELETE FROM favorites
WHERE user_id = ?
  AND place_id = ?;
-- name: FindFavoritesByUserId :many
SELECT place_id
FROM favorites
WHERE user_id = ?;