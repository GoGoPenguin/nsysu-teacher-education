
-- +migrate Up
ALTER TABLE `student_course` MODIFY `review` VARCHAR(300);
ALTER TABLE `student_course` MODIFY `comment` VARCHAR(300);
ALTER TABLE `student_service_learning` MODIFY `review` VARCHAR(300);
ALTER TABLE `student_service_learning` MODIFY `comment` VARCHAR(300);
-- +migrate Down
ALTER TABLE `student_course` MODIFY `review` VARCHAR(150);
ALTER TABLE `student_course` MODIFY `comment` VARCHAR(150);
ALTER TABLE `student_service_learning` MODIFY `review` VARCHAR(36);
ALTER TABLE `student_service_learning` MODIFY `comment` VARCHAR(150);
