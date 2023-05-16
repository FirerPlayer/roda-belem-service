-- name: CreateUser :execute
INSERT INTO users (
    id,
    email,
    avatar,
    username,
    password,
    points,
    missions,
    createdAt,
    updatedAt
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateUserById :execute
UPDATE users
SET ?
WHERE id = ?;
-- name: DeleteUserById :execute
DELETE FROM users
WHERE id = ?;
-- name: UpdateUserPointsByUserId :execute
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