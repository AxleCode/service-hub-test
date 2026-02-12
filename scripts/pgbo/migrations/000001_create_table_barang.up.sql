CREATE TABLE IF NOT EXISTS barang (
    guid CHARACTER VARYING,
    kode_barang VARCHAR(50) NOT NULL UNIQUE,
    nama_barang VARCHAR(150) NOT NULL,
    kategori VARCHAR(150) NOT NULL,
    deskripsi TEXT,
    harga NUMERIC(15,2) NOT NULL DEFAULT 0,
    is_deleted VARCHAR(50) NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT barang_pkey PRIMARY KEY(guid)
);

Insert INTO barang (guid, kode_barang, nama_barang, kategori, deskripsi, harga, is_deleted, created_at, updated_at)
VALUES
(gen_random_uuid(), 'BRG001', 'Indomie Goreng', 'makanan', 'Mie instan goreng', 3500, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG002', 'Indomie Soto', 'makanan', 'Mie instan rasa soto', 3500, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG003', 'Teh Botol Sosro', 'minuman', 'Teh dalam kemasan botol', 5000, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG004', 'Aqua 600ml', 'minuman', 'Air mineral botol', 5000, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG005', 'Beras 5kg', 'sembako', 'Beras premium 5kg', 75000, FALSE, NOW(), NOW()), 
(gen_random_uuid(), 'BRG006', 'Gula Pasir 1kg', 'sembako', 'Gula pasir putih', 15000, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG007', 'Kopi Kapal Api', 'minuman', 'Kopi bubuk hitam', 15000, FALSE, NOW(), NOW()),
(gen_random_uuid(), 'BRG008', 'Roti Tawar', 'makanan', 'Roti tawar putih', 12000, FALSE, NOW(), NOW());