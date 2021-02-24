
-- +migrate Up
ALTER TABLE `service_learning` ADD `show` BOOLEAN COMMENT '是否顯示在前台' AFTER `end`;
-- +migrate Down
ALTER TABLE `service_learning` DROP COLUMN `show`;