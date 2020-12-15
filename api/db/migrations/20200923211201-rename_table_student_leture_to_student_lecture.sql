
-- +migrate Up
RENAME TABLE student_leture TO student_lecture;
-- +migrate Down
RENAME TABLE student_lecture TO student_leture;