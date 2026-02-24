-- name: InsertAuthentication :one
INSERT INTO authentication_schema.authentication(guid, user_guid, username, password, salt, status, created_by)
VALUES (@guid, @user_guid, @username, @password, @salt, @status, @created_by)
RETURNING authentication.*;

-- name: GetAuthenticationIAMByID :one
SELECT a.guid,
        a.id,
        a.user_guid,
        a.username      AS username,
        a.password,
        a.salt,
        a.status,
        u.email         AS email,
        u.phone_number  AS phone_number,
        a.fcm_token,
        a.forgot_password_token,
        a.forgot_password_expiry,
        a.created_at,
        a.created_by,
        a.updated_at,
        a.updated_by
FROM authentication_schema.authentication a
LEFT JOIN 
    authentication_schema.user u ON a.user_guid = u.guid
WHERE a.guid = @guid
    AND a.status != 'deleted'
LIMIT 1;

-- name: GetAuthenticationByID :one
SELECT a.guid,
        a.id,
        a.user_guid,
        u.name as username,
        u.email as email,
        u.phone_number as phone_number,
        a.password,
        a.salt,
        a.status,
        a.fcm_token,
        a.forgot_password_token,
        a.forgot_password_expiry,
        a.created_at,
        a.created_by,
        a.updated_at,
        a.updated_by
FROM authentication_schema.authentication a
LEFT JOIN
    authentication_schema.user u ON a.user_guid = u.guid
WHERE a.guid = @guid
    AND a.status != 'deleted'
LIMIT 1;

-- name: GetAuthenticationByUsername :one
SELECT a.guid,
        a.id,
        a.user_guid,
        u.name as username,
        u.email as email,
        u.phone_number as phone_number,
        a.password,
        a.salt,
        a.status,
        a.fcm_token,
        a.forgot_password_token,
        a.forgot_password_expiry,
        a.created_at,
        a.created_by,
        a.updated_at,
        a.updated_by
FROM authentication_schema.authentication a
LEFT JOIN
    authentication_schema.user u ON a.user_guid = u.guid
WHERE a.username = @username
    AND a.status != 'deleted'
LIMIT 1;

-- name: RecordAuthenticationLastLogin :exec
UPDATE authentication_schema.authentication
SET last_login = (now() at time zone 'UTC')::TIMESTAMP
WHERE guid = @guid
  AND status = 'active';

-- name: UpdateFCMToken :exec
UPDATE authentication_schema.authentication
SET fcm_token = @fcm_token
WHERE guid = @guid
  AND status = 'active';