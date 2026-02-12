-- name: InsertInventory :one
INSERT INTO inventory (
    guid,
    barang_id,
    jumlah,
    keterangan,
    status,
    is_deleted,
    created_at,
    updated_at
    )
VALUES (
    @guid,
    @barang_id,
    @jumlah,
    @keterangan,
    @status,
    FALSE,
    (now() at time zone 'UTC')::TIMESTAMP,
    (now() at time zone 'UTC')::TIMESTAMP
)
RETURNING inventory.*;

-- name: GetOneInventory :one
SELECT 
    guid,
    barang_id,
    jumlah,
    keterangan,
    status,
    created_at,
    updated_at
FROM inventory
WHERE guid = @guid
  AND is_deleted = FALSE;

-- name: ListInventory :many
SELECT 
    guid,
    barang_id,
    jumlah,
    keterangan,
    status,
    created_at,
    updated_at
FROM inventory
WHERE is_deleted = FALSE
    AND (CASE WHEN @set_barang_id::bool THEN barang_id = @barang_id ELSE TRUE END)
    AND (CASE WHEN @set_jumlah::bool THEN jumlah = @jumlah ELSE TRUE END)
    AND (CASE WHEN @set_status::bool THEN status = @status ELSE TRUE END)
ORDER BY
    (CASE WHEN @order_param = 'created_at ASC' THEN created_at END) ASC,
    (CASE WHEN @order_param = 'created_at DESC' THEN created_at END) DESC,
    (CASE WHEN @order_param = 'jumlah ASC' THEN jumlah END) ASC,
    (CASE WHEN @order_param = 'jumlah DESC' THEN jumlah END) DESC,
    created_at DESC
LIMIT @limit_data
OFFSET @offset_pages;

-- name: CountListInventory :one
SELECT 
    COUNT(*) AS count
FROM inventory
WHERE is_deleted = FALSE
    AND (CASE WHEN @set_barang_id::bool THEN barang_id = @barang_id ELSE TRUE END)
    AND (CASE WHEN @set_jumlah::bool THEN jumlah = @jumlah ELSE TRUE END)
    AND (CASE WHEN @set_status::bool THEN status = @status ELSE TRUE END)
    AND is_deleted = FALSE;