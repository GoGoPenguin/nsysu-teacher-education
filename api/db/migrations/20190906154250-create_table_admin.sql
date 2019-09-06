
-- +migrate Up
CREATE TABLE `admin` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `account` VARCHAR(36) NOT NULL COMMENT '帳號',
    `password` BINARY(120) NOT NULL COMMENT '密碼',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    UNIQUE INDEX (`account`),
    INDEX (`deleted_at`),
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理者';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `admin`;
