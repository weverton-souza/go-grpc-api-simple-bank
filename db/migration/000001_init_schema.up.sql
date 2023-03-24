CREATE TABLE `accounts` (
    `id` BIGINT PRIMARY KEY,
    `owner` varchar(255) NOT NULL,
    `balance` bigint NOT NULL,
    `currency` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `entries` (
    `id` BIGINT PRIMARY KEY,
    `account_id` bigint NOT NULL,
    `amount` bigint NOT NULL COMMENT 'can be negative or positive',
    `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `transfers` (
    `id` BIGINT PRIMARY KEY,
    `from_account_id` bigint NOT NULL,
    `to_account_id` bigint NOT NULL,
    `amount` bigint NOT NULL COMMENT 'must be positive',
    `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);