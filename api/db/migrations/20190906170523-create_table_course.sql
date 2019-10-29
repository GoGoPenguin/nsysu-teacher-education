
-- +migrate Up
CREATE TABLE `course` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `topic` VARCHAR(36) NOT NULL COMMENT '研習主題',
    `information` VARCHAR(36) NOT NULL COMMENT '研習資訊',
    `type` ENUM('A', 'B', 'C') NOT NULL COMMENT '研習類別',
    `start` DATETIME NOT NULL COMMENT '開始時間',
    `end` DATETIME NOT NULL COMMENT '結束時間',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    UNIQUE INDEX (`information`),
    INDEX (`topic`),
    INDEX (`type`),
    INDEX (`start`),
    INDEX (`end`),
    INDEX (`deleted_at`),
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='研習';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `course`;
