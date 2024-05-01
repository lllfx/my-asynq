DROP TABLE `task`;

CREATE TABLE `task`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'task ID',
    `task_type`    varchar(45)   NOT NULL COMMENT 'task Type',
    `task_payload` varchar(4096) NOT NULL COMMENT 'task Payload',
    `task_id`      varchar(4096) NOT NULL COMMENT 'task id',
    `task_info`    varchar(4096) NOT NULL COMMENT 'task info',
    `task_status`  int           NOT NULL COMMENT 'task status',
    `err_msg`      varchar(4096) NOT NULL COMMENT 'err msg',
    `create_at`    datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at`    datetime DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;