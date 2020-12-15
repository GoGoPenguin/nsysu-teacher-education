-- +migrate Up
ALTER TABLE `subject_group` CHANGE `leture_type_id` `lecture_type_id` INT;
-- +migrate Down
ALTER TABLE `subject_group` CHANGE `lecture_type_id` `leture_type_id` INT;