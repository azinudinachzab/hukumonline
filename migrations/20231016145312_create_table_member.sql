-- +goose Up
CREATE TABLE IF NOT EXISTS member (
    id          bigint UNSIGNED                     NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name  varchar(255)                        NOT NULL,
    last_name   varchar(255)                        NOT NULL,
    email       varchar(255)                        NOT NULL,
    created_at  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS member;
