DROP USER IF EXISTS 'apkAdmin'; 
CREATE USER 'apkAdmin'@'%'; 
CREATE DATABASE IF NOT EXISTS apkAdmin; 
GRANT ALL ON apkAdmin.* TO 'apkAdmin'@'%' IDENTIFIED BY '123456';

## create table login_user
USE apkAdmin;
CREATE TABLE IF NOT EXISTS  `login_user`(
	`user_id` int unsigned not null AUTO_INCREMENT,
	`username` varchar(100) not null default '',
	`password` varchar(50) not null default '',
	PRIMARY KEY(`user_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

## create tables apk_channel
USE apkAdmin;
CREATE TABLE IF NOT EXISTS `apk_channel`(
	`channelId` int unsigned not null default 0,
	`channel` varchar(50) not null default '',
	PRIMARY KEY(`channelId`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

USE apkAdmin;
CREATE TABLE IF NOT EXISTS `packing_jobs`(
	`job_id` int unsigned not null AUTO_INCREMENT,
	`apk_version` varchar(20) not null default '',
	`apk_channel` varchar(1000) not null default '',
	`creator_id` int unsigned not null default 0,
	`create_time` int unsigned not null default 0,
	`status`  smallint  unsigned not null default 0,
	`download_url` varchar(100) not null default '',
	PRIMARY KEY(`job_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;