-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2)
RETURNING *;


-- name: CreateScramble :one
INSERT INTO scrambles (user_id, time, scramble)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetScrambles :many
SELECT * FROM scrambles
ORDER BY id;

-- name: GetScramblesByUser :many
SELECT * FROM scrambles
WHERE user_id = $1
ORDER BY id;

-- name: UpdateScramble :one
UPDATE scrambles
SET time = $2,
    scramble = $3,
    updated_on = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteScramble :one
DELETE FROM scrambles
WHERE id = $1
RETURNING *;