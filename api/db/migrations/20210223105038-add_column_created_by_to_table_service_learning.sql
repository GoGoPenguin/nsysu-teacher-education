
-- +migrate Up
ALTER TABLE `service_learning` ADD `created_by` INT COMMENT '創建者' AFTER `created_at`;
ALTER TABLE `service_learning` ADD FOREIGN KEY (`created_by`) REFERENCES `student`(`id`) ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE `service_learning` DROP FOREIGN KEY `service_learning_ibfk_1`;
ALTER TABLE `service_learning` DROP COLUMN `created_by`;