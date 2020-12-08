CREATE TABLE `ecust_book`
(
    `id`          bigint unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `author`      varchar(255)    NOT NULL DEFAULT '' COMMENT '作者',
    `title`       varchar(255)    NOT NULL DEFAULT '' COMMENT '标题',
    `press`       varchar(255)    NOT NULL DEFAULT '' COMMENT '出版社',
    `year`        integer         NOT NULL DEFAULT 0 COMMENT '年份',
    `book_id`     varchar(255)    NOT NULL DEFAULT '' COMMENT '索书号',
    `isbn`        varchar(255)    NOT NULL DEFAULT '' COMMENT 'ISBN',
    `Institution` varchar(255)    NOT NULL DEFAULT '' COMMENT '作者所在学院机构',
    `status`      integer         NOT NULL DEFAULT false COMMENT '是否已删除',
    `created_at`  timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    UNIQUE KEY `uniq_book_id` (`book_id`),
    UNIQUE KEY `uniq_isbn` (`isbn`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='著作信息';

CREATE TABLE `ecust_admin`
(
    `id`         bigint unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `username`   varchar(255)    NOT NULL DEFAULT '' COMMENT '用户名',
    `password`   varchar(255)    NOT NULL DEFAULT '' COMMENT '密码',
    `status`     integer         NOT NULL DEFAULT false COMMENT '是否已删除',
    `created_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    UNIQUE KEY `uniq_username` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='管理员';
