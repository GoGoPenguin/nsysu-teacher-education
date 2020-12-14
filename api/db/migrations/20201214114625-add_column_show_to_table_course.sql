
-- +migrate Up
ALTER TABLE `course` ADD `show` BOOLEAN COMMENT '是否顯示在前台' AFTER `type`;
-- +migrate Down
ALTER TABLE `course` DROP COLUMN `show`;