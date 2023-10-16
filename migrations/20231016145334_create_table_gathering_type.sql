-- +goose Up
CREATE TABLE IF NOT EXISTS gathering_type (
    id              int UNSIGNED                        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name            varchar(255)                        NOT NULL,
    created_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS gathering_type;
