-- +migrate Up
ALTER TABLE `lecture_type` CHANGE `leture_category_id` `lecture_category_id` INT;
-- +migrate Down
ALTER TABLE `lecture_type` CHANGE `lecture_category_id` `leture_category_id` INT;