CREATE TABLE `mixinsnapshot` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `snapshot_id` VARCHAR(36) NOT NULL DEFAULT '',
  `trace_id` VARCHAR(36) NOT NULL DEFAULT '',
  `asset_id` VARCHAR(36) NOT NULL DEFAULT '',
  `opponent_id` VARCHAR(36) NOT NULL DEFAULT '',
  `amount` DECIMAL(32,8) NOT NULL DEFAULT '0.00000000',
  `memo` VARCHAR(255) NOT NULL DEFAULT '',
  `type` VARCHAR(255) NOT NULL DEFAULT '',
  `opening_balance` DECIMAL(32,8) NOT NULL DEFAULT '0.00000000',
  `closing_balance` DECIMAL(32,8) NOT NULL DEFAULT '0.00000000',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_trace_id` (`trace_id`),
  INDEX `idx_asset_id` (`asset_id`),
  INDEX `idx_opponent_id` (`opponent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;