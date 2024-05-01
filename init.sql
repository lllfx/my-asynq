DROP TABLE `task`;

CREATE TABLE `task`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `task_type`    varchar(45)   NOT NULL DEFAULT '' COMMENT 'task Type',
    `task_payload` varchar(4096) NOT NULL DEFAULT '' COMMENT 'task Payload',
    `task_id`      varchar(4096) NOT NULL DEFAULT '' COMMENT 'task id',
    `task_info`    varchar(4096) NOT NULL DEFAULT '' COMMENT 'task info',
    `task_status`  int           NOT NULL DEFAULT 0 COMMENT 'task status',
    `err_msg`      varchar(4096) NOT NULL DEFAULT '' COMMENT 'err msg',
    `create_at`    datetime               DEFAULT NULL COMMENT 'Created Time',
    `update_at`    datetime               DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;