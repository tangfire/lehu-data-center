create database if not exists lehu_data_collect_data;

use lehu_data_collect_data;

create table `message_producer_record` (
    `id` bigint AUTO_INCREMENT not null,
    `message_parent_trace_id` bigint default null comment '消息的父级链路id',
    `message_trace_id` bigint DEFAULT NULL COMMENT '消息的链路id',
    `message_id` bigint not null comment '消息id',
    `message_content` text not null comment '消息内容',
    `message_send_exception` text  COMMENT '消息发送失败的异常消息',
    `message_send_status` int default '1' comment '消息发送状态 1-未发送 -1-发送失败 3-发送成功',
    `reconciliation_status` int default '1' comment '消息对账状态 1-未对账 -1-对账完成有问题 2-对账完成没问题 3-对账有问题处理完毕',
    `send_time` datetime default null comment '消息发送时间',
    `status` int default '1' comment '状态 1-启用 0-禁用',
    `update_time` datetime default null comment '更新时间',
    `create_time` datetime default null comment '创建时间',
    primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='消息发送记录表';

CREATE TABLE `message_consumer_record` (
                                             `id` bigint AUTO_INCREMENT NOT NULL,
                                             `message_parent_trace_id` bigint DEFAULT NULL COMMENT '消息的父级链路id',
                                             `message_trace_id` bigint DEFAULT NULL COMMENT '消息的链路id',
                                             `message_id` bigint NOT NULL COMMENT '消息id',
                                             `message_content` text  COMMENT '消息内容',
                                             `message_consumer_exception` text COMMENT '消息消费失败的异常信息',
                                             `message_consumer_status` int DEFAULT '1' COMMENT '消息消费状态 1:未消费 -1:消费失败 2:消费成功',
                                             `message_consumer_count` int NOT NULL DEFAULT '1' COMMENT '消息的消费次数',
                                             `reconciliation_status` int DEFAULT '1' COMMENT '消息对账状态 1:未对账 -1:对账完成有问题 2:对账完成没有问题 3:对账有问题处理完毕',
                                             `consumer_time` datetime DEFAULT NULL COMMENT '消息发送时间',
                                             `status` int DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
                                             `update_time` datetime DEFAULT NULL COMMENT '编辑时间',
                                             `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='消息消费记录表';


CREATE TABLE `video_business_message_consumer_record` (
                                                            `id` bigint auto_increment NOT NULL,
                                                            `message_consumer_record_id` bigint DEFAULT NULL COMMENT '消息消费记录id',
                                                            `video_dimension_type` int DEFAULT '1' COMMENT '视频维度分类，1-父级视频分类 2-视频分类 3-视频本身',
                                                            `date_type` int DEFAULT '1' COMMENT '日期类型',
                                                            `status` int DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
                                                            `update_time` datetime DEFAULT NULL COMMENT '编辑时间',
                                                            `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='视频维度的消息消费记录表';


CREATE TABLE `video_business_message_producer_record` (
                                                            `id` bigint auto_increment NOT NULL,
                                                            `message_producer_record_id` bigint DEFAULT NULL COMMENT '消息发送记录id',
                                                            `video_dimension_type` int DEFAULT '1' COMMENT '视频维度分类，1-父级视频分类 2-视频分类 3-视频本身',
                                                            `date_type` int DEFAULT '1' COMMENT '日期类型',
                                                            `status` int DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
                                                            `update_time` datetime DEFAULT NULL COMMENT '编辑时间',
                                                            `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='视频维度的消息发送记录表';


CREATE TABLE `rule` (
                          `id` bigint auto_increment NOT NULL COMMENT 'id',
                          `rule_describe` varchar(256) CHARACTER SET utf8mb4  DEFAULT NULL COMMENT '规则描述',
                          `rule_name` varchar(256) CHARACTER SET utf8mb4  DEFAULT NULL COMMENT '规则名字',
                          `rule_type` tinyint NOT NULL DEFAULT '1' COMMENT '规则类型 1收集 2 查询',
                          `rule_version_id` bigint NOT NULL COMMENT '如果需要规则变更，直接创建一个新规则，将版本号递增',
                          `status` int DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
                          `update_time` datetime DEFAULT NULL COMMENT '编辑时间',
                          `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='规则表';

INSERT INTO `rule` VALUES (1,'视频数据','video_data',1,1,1,'2025-07-31 10:01:00','2025-07-31 10:01:00');



CREATE TABLE `dimension_gather` (
                                      `id` bigint auto_increment NOT NULL COMMENT 'id',
                                      `rule_id` bigint NOT NULL COMMENT '规则id',
                                      `collect_type` tinyint NOT NULL DEFAULT '1' COMMENT '收集的方式 1：sql查询 2：调用接口',
                                      `collect_detail` text NOT NULL COMMENT ' 具体的收集实现，如果是sql，就是sql脚本。如果是接口，就是url',
                                      `collect_source_name` varchar(128) NOT NULL COMMENT '收集的源头名字，如果是数据库，就是数据源名字',
                                      `entity` varchar(100) DEFAULT NULL COMMENT '表名',
                                      `status` int DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
                                      `update_time` datetime DEFAULT NULL COMMENT '编辑时间',
                                      `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='维度查询表';

INSERT INTO `dimension_gather` VALUES (1,1,1,'SELECT video_id, video_type_id,parent_video_type_id FROM d_video_react where status = 1 and enter_watch_time > :start_time and enter_watch_time < :end_time group by video_id, video_type_id,parent_video_type_id','d_record','d_video_react',1,'2025-07-31 10:01:00','2025-07-31 10:01:00'),(2,1,1,'select id,video_name,video_type_id,video_duration from d_video where status = 1 and video_type_id = :video_type_id','d_record','d_video',1,'2025-07-31 10:01:00','2025-07-31 10:01:00'),(3,1,1,'select id,parent_id,name,category_level from d_video_type where status = 1 and id in ( :video_type_id )','d_record','d_video_type',1,'2025-07-31 10:01:00','2025-07-31 10:01:00');


