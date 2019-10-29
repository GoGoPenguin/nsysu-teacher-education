
-- +migrate Up
CREATE TABLE `student_course` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `student_id` INT NOT NULL COMMENT '學生ID',
    `course_id` INT NOT NULL COMMENT '講座ID',
    `meal` ENUM('meat', 'vegetable') NOT NULL COMMENT '便當',
    `status` ENUM('pass', 'failed', '') NOT NULL DEFAULT '' COMMENT '狀態',
    `review` VARCHAR(150) NOT NULL DEFAULT '' COMMENT '心得',
    `comment` VARCHAR(150) NOT NULL DEFAULT '' COMMENT '備註',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` DATETIME COMMENT '刪除日期',
    INDEX (`status`),
    INDEX (`deleted_at`),
    INDEX (`student_id`),
    INDEX (`course_id`),
    UNIQUE INDEX (`student_id`, `course_id`),
    FOREIGN KEY (`student_id`) REFERENCES `student`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`course_id`) REFERENCES `course`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='學生研習';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `student_course`;
