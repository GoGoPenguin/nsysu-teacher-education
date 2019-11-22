
-- +migrate Up
CREATE TABLE `subject_group` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `leture_type_id` INT NOT NULL COMMENT '課程領域ID',
    `min_credit` TINYINT NOT NULL COMMENT '最低學分數',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`leture_type_id`, `id`),
    INDEX (`deleted_at`),
    FOREIGN KEY (`leture_type_id`) REFERENCES `leture_type`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='科目群組';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `subject_group`;
