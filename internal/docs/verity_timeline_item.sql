create table `item`
(
    `item_id`     bigint unsigned auto_increment,
    `author_id`   bigint unsigned not null comment '作者id',
    `title`       varchar(50) not null comment '标题',
    `type`        tinyint(11) unsigned not null comment '内容类型',
    `cover_url`   text        not null default '' comment '封面图链接',
    `video_url`   text        not null default '' comment '视频链接',
    `picture_url` text        not null default '' comment '图片链接列表，逗号分隔',
    `content`     text        not null default '' comment '文本内容',
    `status`      tinyint unsigned not null default 0 comment '10-删除',
    `created_at`  datetime    not null comment '创建时间',
    `updated_at`  datetime    not null comment '更新时间',
    primary key (`item_id`)
)engine=InnoDB default charset=utf8;