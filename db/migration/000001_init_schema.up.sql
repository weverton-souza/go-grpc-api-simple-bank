CREATE TABLE `account` (
    `id`  varchar(37) primary key,
    `owner` varchar(255) NOT NULL,
    `balance` bigint NOT NULL,
    `currency` varchar(255) NOT NULL,
    `created_at` timestamp default current_timestamp
);

CREATE TABLE `entry` (
    `id` varchar(37) primary key,
    `account_id` varchar(37) NOT NULL,
    `amount` bigint NOT NULL COMMENT 'can be negative or positive',
    `created_at` timestamp default current_timestamp
);

CREATE TABLE `transfer` (
    `id` varchar(37) primary key,
    `from_account_id` varchar(37) NOT NULL,
    `to_account_id` varchar(37) NOT NULL,
    `amount` bigint NOT NULL COMMENT 'must be positive',
    `created_at` timestamp default current_timestamp
);

CREATE INDEX `account_index_0` ON `account` (`owner`);

CREATE INDEX `entry_index_1` ON `entry` (`account_id`);

CREATE INDEX `transfer_index_2` ON `transfer` (`from_account_id`);

CREATE INDEX `transfer_index_3` ON `transfer` (`to_account_id`);

CREATE INDEX `transfer_index_4` ON `transfer` (`from_account_id`, `to_account_id`);

ALTER TABLE `entry` ADD FOREIGN KEY (`account_id`) REFERENCES `account` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`from_account_id`) REFERENCES `account` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`to_account_id`) REFERENCES `account` (`id`);