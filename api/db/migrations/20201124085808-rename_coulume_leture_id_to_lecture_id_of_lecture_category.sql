-- +migrate Up
ALTER TABLE `lecture_category` CHANGE `leture_id` `lecture_id` INT;
-- +migrate Down
ALTER TABLE `lecture_category` CHANGE `lecture_id` `leture_id` INT;