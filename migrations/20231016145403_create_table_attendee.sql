-- +goose Up
CREATE TABLE IF NOT EXISTS attendee (
    member_id       bigint UNSIGNED                     NOT NULL,
    gathering_id    bigint UNSIGNED                     NOT NULL,
    created_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (member_id, gathering_id),
    FOREIGN KEY (member_id) REFERENCES member(id),
    FOREIGN KEY (gathering_id) REFERENCES gathering(id)
);

-- +goose Down
DROP TABLE IF EXISTS attendee;
