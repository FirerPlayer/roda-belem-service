-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(36) NOT NULL PRIMARY KEY,
  `email` varchar(255) NOT NULL UNIQUE,
  `avatar` mediumblob,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `points` int NOT NULL DEFAULT 0,
  `missions` JSON,
  `created_at` datetime,
  `updated_at` datetime
);
CREATE TABLE IF NOT EXISTS `places` (
  `id` varchar(36) NOT NULL PRIMARY KEY,
  `google_place_id` varchar(8000),
  `name` varchar(255),
  `formatted_address` varchar(255),
  `lat` float,
  `lng` float,
  `icon` varchar(255),
  `types` JSON,
  `opening_periods` JSON,
  `photos` JSON,
  `rating` float,
  `created_at` datetime,
  `updated_at` datetime
);
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY,
  `place_id` varchar(255),
  `user_id` varchar(255),
  `content` varchar(255),
  `images` JSON,
  `rating` float,
  `reactions` JSON,
  `accessibility_features` text,
  `created_at` datetime,
  `updated_at` datetime,
  FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `favorites` (
  `place_id` varchar(36),
  `user_id` varchar(36),
  FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
 -- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `favorites`;
DROP TABLE IF EXISTS `reviews`;
DROP TABLE IF EXISTS `places`;
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd