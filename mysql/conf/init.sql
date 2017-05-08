DROP USER IF EXISTS 'apkAdmin'; 
CREATE USER 'apkAdmin'@'%'; 
CREATE DATABASE IF NOT EXISTS apkAdmin; 
USE apkAdmin;
CREATE TABLE `apk_channels`(
	`channelId` int unsigned not null default 0,
	`channel` varchar(50) not null default '',
	PRIMARY KEY(`channelId`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
GRANT ALL ON apkAdmin.* TO 'apkAdmin'@'%' IDENTIFIED BY '123456';
