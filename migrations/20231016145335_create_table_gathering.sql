-- +goose Up
CREATE TABLE IF NOT EXISTS gathering (
    id              bigint UNSIGNED                     NOT NULL AUTO_INCREMENT PRIMARY KEY,
    creator         bigint UNSIGNED                     NOT NULL,
    type            int UNSIGNED                        NOT NULL,
    scheduled_at    timestamp                           NOT NULL,
    name            varchar(255)                        NOT NULL,
    location        varchar(255)                        NOT NULL,
    created_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (creator) REFERENCES member(id),
    FOREIGN KEY (type) REFERENCES gathering_type(id)
);

-- +goose Down
DROP TABLE IF EXISTS gathering;
