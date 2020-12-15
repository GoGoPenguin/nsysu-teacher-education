
-- +migrate Up
RENAME TABLE leture TO lecture;
-- +migrate Down
RENAME TABLE lecture TO leture;