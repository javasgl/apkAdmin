DROP USER IF EXISTS 'apkAdmin'; 
CREATE USER 'apkAdmin'@'%'; 
CREATE DATABASE IF NOT EXISTS apkAdmin; 
GRANT ALL ON apkAdmin.* TO 'apkAdmin'@'%' IDENTIFIED BY '123456';

## create table login_user
USE apkAdmin;
CREATE TABLE `login_user`(
	`user_id` int unsigned not null default 0,
	`username` varchar(100) not null default '',
	`password` varchar(50) not null default '',
	PRIMARY KEY(`user_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
## create tables apk_channel
USE apkAdmin;
CREATE TABLE `apk_channel`(
	`channelId` int unsigned not null default 0,
	`channel` varchar(50) not null default '',
	PRIMARY KEY(`channelId`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

