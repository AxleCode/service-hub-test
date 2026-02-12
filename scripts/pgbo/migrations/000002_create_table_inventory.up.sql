CREATE TABLE IF NOT EXISTS inventory (
    guid CHARACTER VARYING,
    barang_id CHARACTER VARYING NOT NULL,
    jumlah INT NOT NULL CHECK (jumlah > 0),
    keterangan TEXT,
    status VARCHAR(50) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_inventory_barang
        FOREIGN KEY (barang_id)
        REFERENCES barang(guid)
        ON DELETE CASCADE,
        
    CONSTRAINT inventory_pkey PRIMARY KEY(guid)
);

insert INTO inventory (guid, barang_id, jumlah, keterangan, status, is_deleted, created_at, updated_at)
VALUES
(gen_random_uuid(), (SELECT guid FROM barang WHERE kode_barang = 'BRG001'), 100, 'Initial stock for Indomie Goreng in', 'IN', FALSE, NOW(), NOW()),
(gen_random_uuid(), (SELECT guid FROM barang WHERE kode_barang = 'BRG001'), 10, 'Initial stock for Indomie Goreng out', 'OUT', FALSE, NOW(), NOW())