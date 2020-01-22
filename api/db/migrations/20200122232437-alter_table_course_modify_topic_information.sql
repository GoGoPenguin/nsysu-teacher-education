
-- +migrate Up
ALTER TABLE `course` MODIFY `topic` VARCHAR(512);
ALTER TABLE `course` MODIFY `information` VARCHAR(512);
-- +migrate Down
ALTER TABLE `course` MODIFY `topic` VARCHAR(36);
ALTER TABLE `course` MODIFY `information` VARCHAR(36);
