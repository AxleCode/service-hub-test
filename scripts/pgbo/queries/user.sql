-- name: RegisterUser :one
INSERT INTO authentication_schema.user(
    guid, name, phone_number, email, address, created_at
)
VALUES (
    @guid, @name, @phone_number, @email, @address,
    (now() at time zone 'UTC')::TIMESTAMP
)
RETURNING *;
