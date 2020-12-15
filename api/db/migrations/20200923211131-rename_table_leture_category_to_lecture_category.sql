
-- +migrate Up
RENAME TABLE leture_category TO lecture_category;
-- +migrate Down
RENAME TABLE lecture_category TO leture_category;