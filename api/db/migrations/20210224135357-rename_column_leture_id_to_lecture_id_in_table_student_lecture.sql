
-- +migrate Up
ALTER TABLE `student_lecture` CHANGE `leture_id` `lecture_id` INT;
-- +migrate Down
ALTER TABLE `student_lecture` CHANGE `lecture_id` `leture_id` INT;
