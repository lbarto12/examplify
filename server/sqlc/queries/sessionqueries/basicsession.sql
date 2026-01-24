-- name: CreateAccount :one
INSERT INTO user_accounts (normalised_email, user_email, password_hash)
VALUES (UPPER(@user_email), @user_email, @password_hash)
RETURNING *;

-- name: GetUserAccountByEmail :one
SELECT * FROM user_accounts WHERE normalised_email = UPPER(@user_email::varchar);