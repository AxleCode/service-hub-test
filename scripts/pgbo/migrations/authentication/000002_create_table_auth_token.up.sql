CREATE TABLE IF NOT EXISTS authentication_schema.auth_token
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    name                    CHARACTER VARYING           NOT NULL,
    device_id               CHARACTER VARYING           NOT NULL,
    device_type             CHARACTER VARYING           NOT NULL,
    token                   CHARACTER VARYING           NOT NULL,
    token_expired           TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    refresh_token           CHARACTER VARYING           NOT NULL,
    refresh_token_expired   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    is_login                BOOLEAN                     NOT NULL DEFAULT false,
    user_login              CHARACTER VARYING,
    ip_address              CHARACTER VARYING(50),
    created_at              TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at              TIMESTAMP WITHOUT TIME ZONE,
    
    CONSTRAINT token_pkey PRIMARY KEY (name, device_id, device_type)
);

