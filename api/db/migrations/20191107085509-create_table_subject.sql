
-- +migrate Up
CREATE TABLE `subject` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `subject_group_id` INT NOT NULL COMMENT '科目群組ID',
    `name` VARCHAR(150) NOT NULL COMMENT '名稱',
    `credit` INT NOT NULL COMMENT '學分數',
    `compulsory` BOOLEAN NOT NULL COMMENT '必修',
    `status` ENUM('enable', 'disable') NOT NULL DEFAULT 'enable' COMMENT '狀態',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    UNIQUE INDEX (`name`, `subject_group_id`),
    INDEX (`deleted_at`),
    FOREIGN KEY (`subject_group_id`) REFERENCES `subject_group`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='科目';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `subject`;
