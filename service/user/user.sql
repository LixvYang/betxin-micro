CREATE TABLE
    `users` (
        `id` int NOT NULL AUTO_INCREMENT,
        `identity_number` varchar(50) NOT NULL,
        `uid` varchar(36) NOT NULL,
        `full_name` varchar(255) DEFAULT NULL,
        `avatar_url` varchar(255) DEFAULT NULL,
        `session_id` varchar(255) DEFAULT NULL,
        `biography` varchar(255) DEFAULT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        UNIQUE KEY `idx_identity_number` (`identity_number`),
        UNIQUE KEY `idx_uid` (`uid`),
        KEY `idx_full_name` (`full_name`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;