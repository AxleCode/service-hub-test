  CREATE TABLE IF NOT EXISTS authentication_schema.app_key
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name CHARACTER VARYING(100) NOT NULL,
    key CHARACTER VARYING(200)  NOT NULL
);

-- default insert value
INSERT INTO authentication_schema.app_key(name, key)
VALUES ('wit-dev', 'w1t-d3V') ON CONFLICT DO NOTHING;