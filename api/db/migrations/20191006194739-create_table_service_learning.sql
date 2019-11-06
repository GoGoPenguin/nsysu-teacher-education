
-- +migrate Up
CREATE TABLE `service_learning` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `type` ENUM('internship', 'volunteer', 'both') NOT NULL COMMENT '種類',
    `content` VARCHAR(150) NOT NULL COMMENT '服務內容說明',
    `session` VARCHAR(36) NOT NULL COMMENT '時段',
    `hours` INT NOT NULL COMMENT '時數',
    `start` DATETIME NOT NULL COMMENT '開始日期',
    `end` DATETIME NOT NULL COMMENT '結束日期',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`type`),
    INDEX (`start`),
    INDEX (`end`),
    INDEX (`deleted_at`),
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='服務學習';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `service_learning`;
