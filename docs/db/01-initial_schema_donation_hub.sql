CREATE DATABASE donation_hub;
USE donation_hub;

CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` bigint(20) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_roles` (
    `user_id` int(11) NOT NULL,
    `role` enum('admin','donor', 'requester') NOT NULL,
    PRIMARY KEY `user_id_role` (`user_id`,`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `projects` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `target_amount` float NOT NULL,
    `collection_amount` float NOT NULL DEFAULT 0,
    `currency` varchar(255) NOT NULL,
    `status` enum('need_review','approved','completed','rejected') NOT NULL,
    `requester_id` int(11) NOT NULL,
    `due_at` bigint(20) NOT NULL,
    `created_at` bigint(20) NOT NULL,
    `updated_at` bigint(20) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `project_images` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `project_id` int(11) NOT NULL,
    `url` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `donations` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `project_id` int(11) NOT NULL,
    `donor_id` int(11) NOT NULL,
    `message` varchar(255),
    `amount` float NOT NULL,
    `currency` varchar(255) NOT NULL,
    `created_at` bigint(20) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `project_id` (`project_id`),
    INDEX `donor_id` (`donor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

