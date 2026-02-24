-- name: InsertBarang :one
INSERT INTO barang (
    guid,
    kode_barang,
    nama_barang,
    deskripsi,
    kategori,
    harga,
    is_deleted,
    created_at,
    updated_at
    )
VALUES (
    @guid,
    @kode_barang,
    @nama_barang,
    @deskripsi,
    @kategori,
    @harga,
    FALSE,
    (now() at time zone 'UTC')::TIMESTAMP,
    (now() at time zone 'UTC')::TIMESTAMP
)
RETURNING barang.*;

-- name: GetOneBarang :one
SELECT 
    guid,
    kode_barang,
    nama_barang,
    kategori,
    deskripsi,
    harga,
    created_at,
    updated_at
FROM barang
WHERE guid = @guid
  AND is_deleted = FALSE;

-- name: ListBarang :many
SELECT 
    guid,
    kode_barang,
    nama_barang,
    kategori,
    deskripsi,
    harga,
    created_at,
    updated_at
FROM barang
WHERE is_deleted = FALSE
    AND (CASE WHEN @set_kategori::bool THEN LOWER(kategori) LIKE LOWER('%' || @kategori || '%')  ELSE TRUE END)
    AND (CASE WHEN @set_kode_barang::bool THEN LOWER(kode_barang) LIKE LOWER('%' || @kode_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_nama_barang::bool THEN LOWER(nama_barang) LIKE LOWER('%' || @nama_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_guid::bool THEN LOWER(guid) = LOWER(@guid) ELSE TRUE END)
ORDER BY
    (CASE WHEN @order_param = 'created_at ASC' THEN created_at END) ASC,
    (CASE WHEN @order_param = 'created_at DESC' THEN created_at END) DESC,
    (CASE WHEN @order_param = 'harga ASC' THEN harga END) ASC,
    (CASE WHEN @order_param = 'harga DESC' THEN harga END) DESC,
    created_at DESC
LIMIT @limit_data
OFFSET @offset_pages;

-- name: CountListBarang :one
SELECT COUNT(*) AS total_data
FROM barang
WHERE is_deleted = FALSE
    AND (CASE WHEN @set_kategori::bool THEN LOWER(kategori) LIKE LOWER('%' || @kategori || '%')  ELSE TRUE END)
    AND (CASE WHEN @set_kode_barang::bool THEN LOWER(kode_barang) LIKE LOWER('%' || @kode_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_nama_barang::bool THEN LOWER(nama_barang) LIKE LOWER('%' || @nama_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_guid::bool THEN LOWER(guid) = LOWER(@guid) ELSE TRUE END)
    ;

-- name: UpdateBarang :one
UPDATE barang
SET
    kode_barang = @kode_barang,
    nama_barang = @nama_barang,
    kategori = @kategori,
    deskripsi = @deskripsi,
    harga = @harga,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP
WHERE guid = @guid
  AND is_deleted = FALSE
RETURNING barang.*;

-- name: UpdateStatusBarang :one
UPDATE barang
SET
    is_deleted = TRUE,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP
WHERE guid = @guid
  AND is_deleted = FALSE
RETURNING barang.*;