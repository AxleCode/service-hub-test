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
    AND (CASE WHEN @set_barang_id::bool THEN LOWER(barang_id) = LOWER(@barang_id) ELSE TRUE END)
    AND (CASE WHEN @set_guid::bool THEN LOWER(guid) = LOWER(@guid) ELSE TRUE END)
    AND (CASE WHEN @set_status::bool THEN LOWER(status) LIKE LOWER('%' || @status || '%') ELSE TRUE END)
ORDER BY
    (CASE WHEN @order_param = 'created_at ASC' THEN created_at END) ASC,
    (CASE WHEN @order_param = 'created_at DESC' THEN created_at END) DESC,
    created_at DESC
LIMIT @limit_data
OFFSET @offset_pages;

-- name: CountListInventory :one
SELECT 
    COUNT(*) AS count
FROM inventory
WHERE is_deleted = FALSE
    AND (CASE WHEN @set_barang_id::bool THEN LOWER(barang_id) = LOWER(@barang_id) ELSE TRUE END)
    AND (CASE WHEN @set_guid::bool THEN LOWER(guid) = LOWER(@guid) ELSE TRUE END)
    AND (CASE WHEN @set_keterangan::bool THEN LOWER(keterangan) LIKE LOWER('%' || @keterangan || '%') ELSE TRUE END)
    AND (CASE WHEN @set_status::bool THEN LOWER(status) LIKE LOWER('%' || @status || '%') ELSE TRUE END)
;

-- name: UpdateInventory :one
UPDATE inventory
SET
    barang_id = COALESCE(@barang_id, barang_id),
    jumlah = COALESCE(@jumlah, jumlah),
    keterangan = COALESCE(@keterangan, keterangan),
    status = COALESCE(@status, status),
    updated_at = (now() at time zone 'UTC')::TIMESTAMP
WHERE guid = @guid
  AND is_deleted = FALSE
RETURNING inventory.*;

-- name: UpdateStatusInventory :one
UPDATE inventory
SET
    is_deleted = TRUE,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP
WHERE guid = @guid
  AND is_deleted = FALSE
RETURNING inventory.*;

-- name: CountInventoryStokByBarangID :one
SELECT 
    b.guid AS barang_id,
    b.kode_barang,
    b.nama_barang,
    b.kategori,
    COALESCE((
        SELECT SUM(
            CASE 
                WHEN i.status = 'IN'  THEN i.jumlah
                WHEN i.status = 'OUT' THEN -i.jumlah
                ELSE 0
            END
        )
        FROM inventory i
        WHERE i.barang_id = b.guid
          AND i.is_deleted = FALSE
    ),0)::int AS total_stok
FROM barang b
WHERE b.guid = @barang_id;

-- name: ListCountAllInventoryEachProduct :many
SELECT 
    b.guid AS barang_id,
    b.kode_barang,
    b.nama_barang,
    b.kategori,
    b.deskripsi,
    b.harga,
    COALESCE(
        SUM(
            CASE 
                WHEN i.status = 'IN'  THEN i.jumlah
                WHEN i.status = 'OUT' THEN -i.jumlah
                ELSE 0
            END
        ), 
    0)::int AS total_stok
FROM barang AS b
LEFT JOIN inventory AS i 
    ON i.barang_id = b.guid
    AND i.is_deleted = FALSE
WHERE 
    b.is_deleted = FALSE
    AND (CASE WHEN @set_nama_barang::bool THEN LOWER(nama_barang) LIKE LOWER('%' || @nama_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_kategori::bool THEN LOWER(kategori) LIKE LOWER('%' || @kategori || '%') ELSE TRUE END)

GROUP BY 
    b.guid,
    b.kode_barang,
    b.nama_barang,
    b.kategori,
    b.deskripsi,
    b.harga
ORDER BY
    CASE WHEN @order_param = 'nama_barang ASC'  THEN b.nama_barang END ASC,
    CASE WHEN @order_param = 'nama_barang DESC' THEN b.nama_barang END DESC,
    b.nama_barang ASC
LIMIT @limit_data
OFFSET @offset_pages;

-- name: CountListCountAllInventoryEachProduct :one
SELECT 
    COUNT(*) AS total_data
FROM barang AS b
WHERE 
    b.is_deleted = FALSE
    AND (CASE WHEN @set_nama_barang::bool THEN LOWER(nama_barang) LIKE LOWER('%' || @nama_barang || '%') ELSE TRUE END)
    AND (CASE WHEN @set_kategori::bool THEN LOWER(kategori) LIKE LOWER('%' || @kategori || '%') ELSE TRUE END);

-- name: StockItemsByCategory :many
SELECT 
    b.kategori,
    COALESCE(
        SUM(
            CASE 
                WHEN i.status = 'IN'  THEN i.jumlah
                WHEN i.status = 'OUT' THEN -i.jumlah
                ELSE 0
            END
        ), 
    0)::int AS total_stok
FROM barang AS b
LEFT JOIN inventory AS i 
    ON i.barang_id = b.guid
    AND i.is_deleted = FALSE
WHERE 
    b.is_deleted = FALSE
    AND (CASE WHEN @set_kategori::bool 
        THEN LOWER(b.kategori) LIKE LOWER('%' || @kategori || '%') 
        ELSE TRUE 
    END)
GROUP BY 
    b.kategori
ORDER BY
    CASE WHEN @order_param = 'kategori ASC'  THEN b.kategori END ASC,
    CASE WHEN @order_param = 'kategori DESC' THEN b.kategori END DESC,
    b.kategori ASC
LIMIT @limit_data
OFFSET @offset_pages;

-- name: CountStockItemsByCategory :one
SELECT 
    COUNT(DISTINCT b.kategori) AS total_data
FROM barang AS b
WHERE 
    b.is_deleted = FALSE
    AND (CASE WHEN @set_kategori::bool THEN LOWER(b.kategori) LIKE LOWER('%' || @kategori || '%') ELSE TRUE END);