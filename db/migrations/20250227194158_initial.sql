-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`user_id` integer NULL PRIMARY KEY AUTOINCREMENT, `first_name` text NOT NULL, `last_name` text NOT NULL, `email` text NOT NULL, `password` text NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP));
-- Drop "users" table without copying rows (no columns)
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_email" to table: "users"
CREATE UNIQUE INDEX `users_email` ON `users` (`email`);
-- Create "tenants" table
CREATE TABLE `tenants` (`tenant_id` integer NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
