-- Create "users" table
CREATE TABLE `users` (
  `user_id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `pid` text NOT NULL,
  `first_name` text NOT NULL,
  `last_name` text NOT NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Create "tenants" table
CREATE TABLE `tenants` (
  `tenant_id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL
);
