
-- +migrate Up
ALTER TABLE `student_subject` CHANGE `student_leture_id` `student_lecture_id` INT;
-- +migrate Down
ALTER TABLE `student_subject` CHANGE `student_lecture_id` `student_leture_id` INT;
