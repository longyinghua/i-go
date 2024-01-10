CREATE TABLE `book` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `title` varchar(128) NOT NULL COMMENT '书籍名称',
                        `author` varchar(128) NOT NULL COMMENT '作者',
                        `price` int NOT NULL DEFAULT '0' COMMENT '价格',
                        `publish_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '出版日期',
                        `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=602 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='书籍表';