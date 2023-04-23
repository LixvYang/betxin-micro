CREATE TABLE
    `topic` (
        `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
        `tid` VARCHAR(36) NOT NULL DEFAULT '',
        `cid` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
        `title` VARCHAR(50) NOT NULL DEFAULT '',
        `intro` VARCHAR(255) NOT NULL DEFAULT '',
        `content` VARCHAR(255) NOT NULL DEFAULT '',
        `yes_ratio` DECIMAL(5, 2) NOT NULL DEFAULT 50.00,
        `no_ratio` DECIMAL(5, 2) NOT NULL DEFAULT 50.00,
        `yes_price` DECIMAL(16, 8) NOT NULL DEFAULT 0.00000000,
        `no_price` DECIMAL(16, 8) NOT NULL DEFAULT 0.00000000,
        `total_price` DECIMAL(32, 8) NOT NULL DEFAULT 0.00000000,
        `collect_count` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
        `read_count` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
        `img_url` VARCHAR(255) NOT NULL DEFAULT '',
        `is_stop` TINYINT UNSIGNED NOT NULL DEFAULT '0',
        `refund_end_time` TIMESTAMP DEFAULT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (`id`),
        INDEX `idx_cid` (`cid`),
        UNIQUE `idx_tid` (`tid`),
        UNIQUE `title_intro_content_index` (`title`, `intro`, `content`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE
    `category` (
        `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
        `category_name` VARCHAR(30) NOT NULL DEFAULT '',
        PRIMARY KEY (`id`)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;