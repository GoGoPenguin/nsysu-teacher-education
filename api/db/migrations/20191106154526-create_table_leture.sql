
-- +migrate Up
CREATE TABLE `leture` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150) NOT NULL COMMENT '名稱',
    `min_credit` TINYINT NOT NULL COMMENT '最低學分數',
    `comment` VARCHAR(300) NOT NULL COMMENT '備註',
    `status` ENUM('enable', 'disable') NOT NULL DEFAULT 'enable' COMMENT '狀態',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    UNIQUE INDEX (`name`),
    INDEX (`deleted_at`),
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='課程';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `leture`;
