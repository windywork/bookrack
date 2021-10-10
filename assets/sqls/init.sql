SET FOREIGN_KEY_CHECKS=0;
drop DATABASE IF EXISTS `bookrick`;
CREATE DATABASE `bookrick`;
USE `bookrick`;

/* 用户表 */
DROP TABLE IF EXISTS `t_users`;
CREATE TABLE `t_users`(
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  `fullname` VARCHAR(120) NOT NULL COMMENT '姓名',
  `username` VARCHAR(120) UNIQUE NOT NULL COMMENT '用户名',
  `phone` VARCHAR(120) NULL COMMENT '手机号',
  `email` VARCHAR(120) NULL COMMENT '邮箱',
  `password` VARCHAR(120) NOT NULL COMMENT '密码',
  `avatar` VARCHAR(120) NULL COMMENT '头像',
  `level` TINYINT(1) DEFAULT 1 COMMENT '用户级别',
  `published` TINYINT(1) DEFAULT 1 COMMENT '是否启动',
  `intro` TEXT NULL COMMENT '备注',
  `create_at` BIGINT(20) NOT NULL COMMENT '创建时间',
  `update_at` BIGINT(20) NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 文章类别 */
DROP TABLE IF EXISTS `t_article_categories`;
CREATE TABLE `t_article_categories` (
  `id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '编号',
  `name` VARCHAR(50) NOT NULL UNIQUE COMMENT '名称',
  `layout` VARCHAR(50) NOT NULL COMMENT '类型',
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `description` TEXT NULL COMMENT '描述',
  `published` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `create_at` BIGINT(20) NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


/* 文章 */
DROP TABLE IF EXISTS `t_articles`;
CREATE TABLE `t_articles` (
  `id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT  COMMENT '编号',
  `name` VARCHAR(40) NOT NULL UNIQUE COMMENT '名称',
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `cover` VARCHAR(500) NULL COMMENT '封面',
  `label` VARCHAR(500) NULL COMMENT '标签',
  `summary` TEXT NULL COMMENT '摘要',
  `description` TEXT NULL COMMENT '描述',
  `category_id` INT(11) NOT NULL COMMENT '分类编号',
  `user_id` INT(11) NOT NULL COMMENT '作者',
  `attachment`  VARCHAR(500) NULL COMMENT '附件',
  `views_count` INT(11) NOT NULL DEFAULT 0 COMMENT '查看数',
  `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `state` smallINT(2) DEFAULT 0 COMMENT '状态',
  `create_at` BIGINT(20) NOT NULL COMMENT '创建时间',
  `update_at` BIGINT(20) NOT NULL COMMENT '更新时间',
  `published` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `recommend` TINYINT(1) DEFAULT 0 COMMENT '是否推荐',
  `deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
  FOREIGN KEY (`user_id`) REFERENCES `t_users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`category_id`) REFERENCES `t_article_categories`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


SET FOREIGN_KEY_CHECKS=1;