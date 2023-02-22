CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `email` varchar(50) UNIQUE NOT NULL,
  `password` varchar(50) NOT NULL,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `groups` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `group_members` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `groupId` int NOT NULL,
  `userId` int NOT NULL,
  `joined_at` datetime DEFAULT (now())
);

CREATE TABLE `group_chat` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `groupId` int NOT NULL,
  `userId` int NOT NULL,
  `message` text NOT NULL,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `friends` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `userId` int NOT NULL,
  `friendId` int NOT NULL,
  `created_at` datetime DEFAULT (now())
);

CREATE INDEX `users_index_0` ON `users` (`name`);

CREATE INDEX `users_index_1` ON `users` (`email`);

CREATE INDEX `groups_index_2` ON `groups` (`name`);

CREATE INDEX `group_members_index_3` ON `group_members` (`groupId`);

CREATE INDEX `group_members_index_4` ON `group_members` (`userId`);

CREATE INDEX `group_chat_index_5` ON `group_chat` (`groupId`);

CREATE INDEX `group_chat_index_6` ON `group_chat` (`userId`);

ALTER TABLE `group_members` ADD FOREIGN KEY (`groupId`) REFERENCES `groups` (`id`);

ALTER TABLE `group_members` ADD FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `group_chat` ADD FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `group_chat` ADD FOREIGN KEY (`groupId`) REFERENCES `groups` (`id`);

ALTER TABLE `friends` ADD FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `friends` ADD FOREIGN KEY (`friendId`) REFERENCES `users` (`id`);

ALTER TABLE `group_chat` ADD FOREIGN KEY (`userId`) REFERENCES `group_chat` (`id`);
