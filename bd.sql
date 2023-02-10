CREATE DATABASE 'go-simple-blog' DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;

USE 'go-simple-blog';

CREATE TABLE `posts` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL DEFAULT 'Empty article :P' COLLATE 'latin1_swedish_ci',
	`content` TEXT NOT NULL COLLATE 'latin1_swedish_ci',
	`author` VARCHAR(50) NOT NULL DEFAULT 'Neil Amnstrong' COLLATE 'latin1_swedish_ci',
	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=4
;
