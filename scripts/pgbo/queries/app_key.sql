-- name: GetAppKeyByName :one
SELECT ak.id,
         ak.name,
         ak.key
FROM authentication_schema.app_key ak
WHERE ak.name = @name;