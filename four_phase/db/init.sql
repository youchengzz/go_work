CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE posts (
	id INT auto_increment NOT NULL,
	title varchar(100) NULL,
	content varchar(100) NULL,
	user_id int NULL,
	create_at TIMESTAMP NULL,
	update_at TIMESTAMP NULL,
	CONSTRAINT posts_pk PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE table comments (
	id INT auto_increment NOT NULL,
	content varchar(100) NULL,
	user_id INT NULL,
	post_id INT NULL,
	create_at timestamp NULL,
	update_at TIMESTAMP NULL,
	CONSTRAINT comments_pk PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `path` varchar(100) DEFAULT NULL,
  `requart_param` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  `ip` varchar(100) DEFAULT NULL,
  `error_message` varchar(100) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;