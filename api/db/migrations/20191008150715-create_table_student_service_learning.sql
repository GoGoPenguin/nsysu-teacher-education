
-- +migrate Up
CREATE TABLE `student_service_learning` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_id` INT NOT NULL COMMENT '學生ID',
    `service_learning_id` INT NOT NULL COMMENT '講座ID',
    `status` ENUM('pass', 'failed', '') NOT NULL DEFAULT '' COMMENT '狀態',
    `review` VARCHAR(36) NOT NULL DEFAULT '' COMMENT '心得',
    `reference` VARCHAR(36) NOT NULL COMMENT '佐證資料',
    `comment` VARCHAR(150) NOT NULL DEFAULT '' COMMENT '備註',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`status`),
    INDEX (`deleted_at`),
    INDEX (`student_id`),
    INDEX (`service_learning_id`),
    UNIQUE INDEX (`student_id`, `service_learning_id`),
    FOREIGN KEY (`student_id`) REFERENCES `student`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`service_learning_id`) REFERENCES `service_learning`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='學生服務學習';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `student_service_learning`;
