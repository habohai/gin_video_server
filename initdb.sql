create table `video_user`( 
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL COMMENT '账号',
    `pwd` varchar(64) DEFAULT '' COMMENT '密码',
    unique key (`user_name`),
    PRIMARY KEY(id) 
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

create table `video_comment`(
    `id` varchar(64) NOT NULL COMMENT '评论的唯一标示',
    `user_id` int(10) DEFAULT '0' COMMENT '创建评论的用户id',
    `user_name` varchar(64) NOT NULL COMMENT '账号',
    `video_id` varchar(64) NOT NULL COMMENT "评论对应的视频id",
    `content` text COMMENT "评论内容",
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (id) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='视频评论';

create table `video_videoinfo`( 
    `id` varchar(64) NOT NULL COMMENT '视频的唯一标示',
    `user_id` int(10) COMMENT "上传视频的用户对应的id",
    `user_name` varchar(64) NOT NULL COMMENT '账号',
    `video_name` text COMMENT "视频名称",
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (id) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='视频信息';

create table `video_session`(
    `id` varchar(64) NOT NULL COMMENT "cookic中存的session",
    `ttl` int(10) unsigned DEFAULT '0' COMMENT 'session的到期时间',
    `user_name` varchar(64) NOT NULL COMMENT '账号',
    PRIMARY KEY (`id`) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户session';

create table `video_delrec`(
    `video_id` varchar(64) NOT NULL COMMENT "要删除的视频对应video_id",
    PRIMARY KEY(`video_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='视频删除记录表';

