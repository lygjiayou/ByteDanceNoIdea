CREATE DATABASE IF NOT EXISTS `noideadouyin` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `noideadouyin`;
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name` varchar(20)        NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` varchar(20)     NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender`    varchar(10)    NOT NULL DEFAULT 1 COMMENT   'male男性female女性',
    `follow_count`          int(10) unsigned NOT NULL  COMMENT '关注总数',
    `follower_count`        int(10) unsigned NOT NULL  COMMENT '粉丝总数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user_info`
VALUES (1, '小明','123','male',1,1),
       (2, '小红','124','female',1,1);

DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info`
(
    `id`      int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `title`     varchar(20) NOT NULL DEFAULT 0 COMMENT '视频标题',
    `name`       varchar(20)        NOT NULL default '' COMMENT '视频作者ID',
    `play_url`     varchar(20)                NOT NULL COMMENT '视频播放地址',
    `cover_url`     varchar(20)                NOT NULL COMMENT '视频封面地址',
    `favorite_count`     int(10)                NOT NULL COMMENT '视频点赞总数',
    `comment_count`     int(10)                NOT NULL COMMENT '视频评论总数',
    `issue_time` datetime           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '投稿时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

INSERT INTO `video_info`
VALUES (1, '原生', 1, 'douyin/feed', '/douyin/favorite', 1, 0, '2022-04-01 13:50:19');

DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频评论ID',
    `video_id`   int(10) unsigned NOT NULL DEFAULT 0 COMMENT '视频ID',
    `commenter_id`     int(10) unsigned NOT NULL DEFAULT 0 COMMENT '评论用户ID',
    `content`     TINYTEXT                NOT NULL COMMENT '评论内容',
    `create_time` DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布日期',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';
INSERT INTO `comment_info`
VALUES (1, 1, 1, '有趣', '2022-04-01 14:50:19');

DROP TABLE IF EXISTS `like_info`;
CREATE TABLE `like_info`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '点赞记录ID',
    `user_id`   int(10) unsigned NOT NULL DEFAULT 0 COMMENT '点赞用户ID',
    `video_id`     int(10) unsigned NOT NULL DEFAULT 0 COMMENT '被点赞视频ID',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='点赞信息表';
INSERT INTO `like_info`
VALUES (1, 2, 1);

DROP TABLE IF EXISTS `follow_info`;
CREATE TABLE `follow_info`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '关注记录ID',
    `follower_id`   int(10) unsigned NOT NULL DEFAULT 0 COMMENT '关注者用户ID',
    `star_id`     int(10) unsigned NOT NULL DEFAULT 0 COMMENT '被关注者用户ID',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='关注信息表';
INSERT INTO `follow_info`
VALUES (1, 1, 2),
       (2, 2, 1);