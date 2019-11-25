
-- +migrate Up
CREATE TABLE `student_subject` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_leture_id` INT NOT NULL COMMENT '學生課程ID',
    `subject_id` INT NOT NULL COMMENT '教育部核定科目ID',
    `name` VARCHAR(150) NOT NULL DEFAULT '' COMMENT '師資生已修習科目名稱',
    `year` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '學年度',
    `semester` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '學期',
    `credit` VARCHAR(10) NOT NULL  DEFAULT ''COMMENT '學分',
    `score` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '成績',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`deleted_at`),
    UNIQUE INDEX (`student_leture_id`, `subject_id`),
    FOREIGN KEY (`student_leture_id`) REFERENCES `student_leture`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`subject_id`) REFERENCES `subject`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='學生科目';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `student_subject`;
