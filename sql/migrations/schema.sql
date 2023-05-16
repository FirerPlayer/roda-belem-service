CREATE TABLE IF NOT EXISTS `users` (
  `id` uuid,
  `email` varchar(255) NOT NULL UNIQUE,
  `avatar` mediumblob,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `points` int NOT NULL DEFAULT 0,
  `missions` JSON,
  `createdAt` datetime,
  `updatedAt` datetime
);
CREATE TABLE IF NOT EXISTS `places` (
  `id` uuid,
  `placeId` varchar(255),
  `name` varchar(255),
  `formatted_address` varchar(255),
  `lat` float,
  `lng` float,
  `icon` varchar(255),
  `types` JSON,
  `opening_periods` JSON,
  `photos` JSON,
  `rating` float,
  `AccessibilityFeatures` JSON
);
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` uuid,
  `placeId` varchar(255),
  `userId` varchar(255),
  `text` varchar(255),
  `images` JSON,
  `rating` float,
  `reactions` JSON,
  `createdAt` datetime,
  `updatedAt` datetime
);
CREATE TABLE IF NOT EXISTS `favorites` (
  `placeId` varchar(255),
  `userId` varchar(255)
);