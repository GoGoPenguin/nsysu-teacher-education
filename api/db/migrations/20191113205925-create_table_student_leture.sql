
-- +migrate Up
CREATE TABLE `student_leture` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_id` INT NOT NULL COMMENT '學生ID',
    `leture_id` INT NOT NULL COMMENT '課程ID',
    `pass` BOOLEAN NOT NULL COMMENT '通過',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`pass`),
    INDEX (`deleted_at`),
    UNIQUE INDEX (`student_id`, `leture_id`),
    FOREIGN KEY (`student_id`) REFERENCES `student`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`leture_id`) REFERENCES `leture`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='學生課程';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `student_leture`;
