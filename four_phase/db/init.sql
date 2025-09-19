CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE test.posts (
	id INT auto_increment NOT NULL,
	title varchar(100) NULL,
	content varchar(100) NULL,
	user_id int NULL,
	create_at TIMESTAMP NULL,
	update_at TIMESTAMP NULL,
	CONSTRAINT posts_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;

CREATE TABLE test.comments (
	id INT auto_increment NOT NULL,
	content varchar(100) NULL,
	user_id INT NULL,
	post_id INT NULL,
	create_at timestamp NULL,
	update_at TIMESTAMP NULL,
	CONSTRAINT comments_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;

