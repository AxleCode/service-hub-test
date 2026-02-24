CREATE TABLE IF NOT EXISTS authentication_schema.authentication
(
    guid                    CHARACTER VARYING, 
    id                      BIGINT  GENERATED ALWAYS AS IDENTITY,
    user_guid               CHARACTER VARYING           UNIQUE,
    username                CHARACTER VARYING           NOT NULL UNIQUE,
    password                CHARACTER VARYING           NOT NULL,
    salt                    CHARACTER VARYING,
    forgot_password_token   CHARACTER VARYING,
    forgot_password_expiry  TIMESTAMP,
    is_active               BOOLEAN                     NOT NULL DEFAULT TRUE,
    last_login              TIMESTAMP WITH TIME ZONE,
    fcm_token               TEXT,
    
    status                  CHARACTER VARYING           NOT NULL DEFAULT('active'),

    created_at              TIMESTAMP WITH TIME ZONE    DEFAULT CURRENT_TIMESTAMP,
    created_by              CHARACTER VARYING           NOT NULL,
    updated_at              TIMESTAMP WITH TIME ZONE,
    updated_by              CHARACTER VARYING,

    CONSTRAINT authentication_pkey PRIMARY KEY(guid)
);