DROP USER IF EXISTS 'apkAdmin'; 
CREATE USER 'apkAdmin'@'%'; 

DROP DATABASE IF EXISTS `apkAdmin`;
CREATE DATABASE IF NOT EXISTS `apkAdmin`;

GRANT ALL ON apkAdmin.* TO 'apkAdmin'@'%' IDENTIFIED BY '123456';

## create table apk_users
USE apkAdmin;
CREATE TABLE IF NOT EXISTS  `apk_users`(
	`user_id` int unsigned not null AUTO_INCREMENT,
	`username` varchar(100) not null default '',
	`password` varchar(50) not null default '',
	PRIMARY KEY(`user_id`),
	UNIQUE KEY `username` (`username`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

## create table apk_channels
USE apkAdmin;
CREATE TABLE IF NOT EXISTS `apk_channels`(
	`channel_id` int unsigned not null AUTO_INCREMENT,
	`channel` varchar(50) not null default '',
	`remark` varchar(100) not null default '',
	`create_time` int unsigned not null default 0,
	PRIMARY KEY(`channel_id`),
	UNIQUE KEY `channel` (`channel`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

## create table apk_jobs
USE apkAdmin;
CREATE TABLE IF NOT EXISTS `apk_jobs`(
	`job_id` int unsigned not null AUTO_INCREMENT,
	`app_id` int unsigned not null default 1,
	`release_type` int unsigned not null default 0,
	`app_name` varchar(100) not null default '',
	`apk_version` varchar(20) not null default '',
	`apk_channel` varchar(1000) not null default '',
	`creator_id` int unsigned not null default 0,
	`create_time` int unsigned not null default 0,
	`status`  smallint  unsigned not null default 0,
	`download_url` varchar(100) not null default '',
	`splash_image` varchar(200) not null default '',
	PRIMARY KEY(`job_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;