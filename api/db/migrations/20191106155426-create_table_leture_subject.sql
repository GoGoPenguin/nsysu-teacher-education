
-- +migrate Up
CREATE TABLE `leture_type` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `leture_category_id` INT NOT NULL COMMENT '課程領域ID',
    `name` VARCHAR(150) NOT NULL COMMENT '名稱',
    `min_credit` INT NOT NULL COMMENT '最低學分數',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`leture_category_id`, `name`),
    INDEX (`deleted_at`),
    FOREIGN KEY (`leture_category_id`) REFERENCES `leture_category`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='課程類別';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `leture_type`;
