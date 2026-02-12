CREATE TABLE IF NOT EXISTS inventory (
    guid CHARACTER VARYING,
    barang_id CHARACTER VARYING NOT NULL,
    stok INT NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_inventory_barang
        FOREIGN KEY (barang_id)
        REFERENCES barang(guid)
        ON DELETE CASCADE,
        
    CONSTRAINT inventory_pkey PRIMARY KEY(guid)
);
