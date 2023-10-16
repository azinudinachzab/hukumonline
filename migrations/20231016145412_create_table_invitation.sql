-- +goose Up
CREATE TABLE IF NOT EXISTS invitation (
    id              bigint UNSIGNED                     NOT NULL AUTO_INCREMENT PRIMARY KEY,
    member_id       bigint UNSIGNED                     NOT NULL,
    gathering_id    bigint UNSIGNED                     NOT NULL,
    status          tinyint   DEFAULT 0                 NOT NULL COMMENT '0 PENDING, 1 ACCEPTED, 2 REJECTED',
    created_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (member_id) REFERENCES member(id),
    FOREIGN KEY (gathering_id) REFERENCES gathering(id)
);

-- +goose Down
DROP TABLE IF EXISTS invitation;
