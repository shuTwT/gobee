-- Create "albums" table
CREATE TABLE `albums` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` text NOT NULL,
  `description` text NULL,
  PRIMARY KEY (`oid`)
);
-- Create "articles" table
CREATE TABLE `articles` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  PRIMARY KEY (`oid`)
);
-- Create "comments" table
CREATE TABLE `comments` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `post_id` integer NOT NULL,
  `page_id` integer NOT NULL,
  `content` text NOT NULL,
  `user_id` integer NOT NULL,
  `status` integer NOT NULL,
  `user_agent` text NULL,
  `ip_address` text NOT NULL,
  `ip_location` text NULL,
  `pinned` bool NOT NULL DEFAULT false,
  PRIMARY KEY (`oid`)
);
-- Create "files" table
CREATE TABLE `files` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` text NOT NULL,
  `path` text NOT NULL,
  `url` text NOT NULL,
  `type` text NOT NULL,
  `size` text NOT NULL,
  PRIMARY KEY (`oid`)
);
-- Create "model_schemas" table
CREATE TABLE `model_schemas` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT
);
-- Create "pages" table
CREATE TABLE `pages` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `description` text NULL,
  PRIMARY KEY (`oid`)
);
-- Create "pay_channels" table
CREATE TABLE `pay_channels` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` text NOT NULL,
  `code` text NOT NULL,
  `type` text NOT NULL,
  `config` text NOT NULL,
  PRIMARY KEY (`oid`)
);
-- Create "pay_orders" table
CREATE TABLE `pay_orders` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `channel_id` text NOT NULL,
  `order_id` text NOT NULL,
  `out_trade_no` text NOT NULL,
  `total_fee` text NOT NULL,
  `subject` text NOT NULL,
  `body` text NOT NULL,
  `notify_url` text NOT NULL,
  `return_url` text NOT NULL,
  `extra` text NOT NULL,
  `pay_url` text NULL,
  `state` text NOT NULL DEFAULT '1',
  `error_msg` text NULL,
  `raw` text NULL,
  PRIMARY KEY (`oid`)
);
-- Create "settings" table
CREATE TABLE `settings` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `key` text NOT NULL,
  `value` text NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`oid`)
);
-- Create "users" table
CREATE TABLE `users` (
  `oid` uuid NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `email` text NOT NULL,
  `email_verified` bool NOT NULL DEFAULT false,
  `name` text NOT NULL DEFAULT 'unknown',
  `phone_number` text NULL,
  `phone_number_verified` bool NOT NULL DEFAULT false,
  `password` text NOT NULL,
  PRIMARY KEY (`oid`)
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX `users_phone_number_key` ON `users` (`phone_number`);
