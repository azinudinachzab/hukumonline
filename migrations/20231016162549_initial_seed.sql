-- +goose Up
INSERT INTO gathering_type (id, name) VALUES (1, 'FORMAL'), (2, 'NON FORMAL'), (3, 'LAINNYA');

-- +goose Down
DELETE FROM gathering_type;
