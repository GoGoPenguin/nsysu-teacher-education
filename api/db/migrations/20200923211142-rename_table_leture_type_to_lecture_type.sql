
-- +migrate Up
RENAME TABLE leture_type TO lecture_type;
-- +migrate Down
RENAME TABLE lecture_type TO leture_type;