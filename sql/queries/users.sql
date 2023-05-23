-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    avatar,
    username,
    password,
    points,
    missions,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: GetAllUsers :many
SELECT *
FROM users;
-- name: UpdateUserById :exec
UPDATE users
SET email = ?,
  avatar = ?,
  username = ?,
  password = ?,
  points = ?,
  missions = ?,
  updated_at = ?
WHERE id = ?;
-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = ?;
-- name: UpdateUserPointsByUserId :exec
UPDATE users
SET points = ?
WHERE id = ?;
-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
-- name: FindUserById :one
SELECT *
FROM users
WHERE id = ?;