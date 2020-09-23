
-- +migrate Up
ALTER TABLE `student_service_learning` ADD `hours` INT AFTER `reference`;
-- +migrate Down
ALTER TABLE `student_service_learning` DROP COLUMN `hours`;