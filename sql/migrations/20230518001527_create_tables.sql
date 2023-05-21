-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
  `id` uuid,
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
  `id` uuid,
  `place_id` varchar(255),
  `name` varchar(255),
  `formatted_address` varchar(255),
  `coordinates` POINT,
  `icon` varchar(255),
  `types` JSON,
  `opening_periods` JSON,
  `photos` JSON,
  `rating` float,
  `accessibility_features` JSON
);
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` uuid,
  `place_id` varchar(255),
  `user_id` varchar(255),
  `text` varchar(255),
  `images` JSON,
  `rating` float,
  `reactions` JSON,
  `created_at` datetime,
  `updated_at` datetime,
  FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `favorites` (
  `place_id` varchar(255),
  `user_id` varchar(255),
  FOREIGN KEY (place_id) REFERENCES places (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd