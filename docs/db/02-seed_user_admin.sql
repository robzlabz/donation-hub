USE donation_hub;

INSERT INTO `users` (`id`, `username`, `email`, `password`, `created_at`)
VALUES (1, 'admin', 'admin@donationhub.com', 'admin123', UNIX_TIMESTAMP());

INSERT INTO `user_roles` (`user_id`, `role`)
VALUES (1, 'admin');
