CREATE DATABASE `admin` DEFAULT CHARACTER SET = `utf8mb4`;

USE `admin`;

CREATE TABLE `access_log` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid` int(11) unsigned NOT NULL COMMENT 'user.id',
    `method` varchar(255) NOT NULL COMMENT 'request method',
    `path` varchar(255) NOT NULL COMMENT 'path',
    `header` json NOT NULL COMMENT 'http headers',
    `query` json NOT NULL COMMENT 'http query',
    `body` json NOT NULL COMMENT 'http body',
    `ip` varchar(32) NOT NULL COMMENT 'client ip',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `group_menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `mid` int(11) unsigned NOT NULL COMMENT 'menu.id',
    `gid` int(11) unsigned NOT NULL COMMENT 'group.id',
   	 PRIMARY KEY (`id`),
    KEY `idx_gid` (`gid`),
    KEY `idx_mid` (`mid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `group_user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid` int(11) unsigned NOT NULL COMMENT 'uid',
    `gid` int(11) unsigned NOT NULL COMMENT 'group.id',
    	PRIMARY KEY (`id`),
    KEY `idx_gid` (`gid`),
    KEY `idx_uid` (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
INSERT INTO `group_user` (`id`, `uid`, `gid`)
VALUES
	(1, 1, 1);

CREATE TABLE `group` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '群组名称',
    `status` tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '状态（1-正常，0-禁用）',
    `remark` varchar(200) NOT NULL DEFAULT '',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO
    `group` (`id`, `name`, `remark`)
VALUES
    (1, 'admin', 'super admin have all access');

CREATE TABLE `menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '菜单名称',
    `uri` varchar(100) DEFAULT '' COMMENT 'URI地址',
    `level` tinyint(2) unsigned DEFAULT '1' COMMENT '节点等级',
    `pid` int(11) unsigned DEFAULT '0' COMMENT '父节点菜单ID',
    `icon` varchar(100) DEFAULT '' COMMENT '菜单图标',
    `status` tinyint(1) unsigned DEFAULT '1' COMMENT '状态（0-禁用，1-正常）',
    `sort` int(11) unsigned DEFAULT '999' COMMENT '排序',
    `remark` varchar(255) DEFAULT '' COMMENT '备注',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_pid` (`pid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO `menu` (`id`, `name`, `uri`, `level`, `pid`, `icon`, `status`, `sort`, `remark`)
VALUES
	(1, 'Home', '/home', 1, 0, 'dashboard', 1, 100, ''),
	(2, 'Management', '/panel', 1, 0, 'config', 1, 900, ''),
	(3, 'Dashboard', '/home/index', 2, 1, 'dashboard', 1, 101, ''),
	(4, 'User', '/panel/users', 2, 2, 'user', 1, 901, ''),
	(5, 'Groups', '/panel/groups', 2, 2, 'group', 1, 902, ''),
	(6, 'Menus', '/panel/menus', 2, 2, 'tree', 1, 903, '');



CREATE TABLE `user` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `login_time` int unsigned NOT NULL,
  `pwd_wrong` tinyint unsigned NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `user` (`uid`, `name`, `password`, `salt`, `real_name`, `mobile`, `email`, `desc`, `login_time`, `pwd_wrong`, `status`)
VALUES
	(1, 'admin', '6bd26abb93950168b41fb41d75f69f86', 'hM9MUs', '超管', '18888888888', 'admin@admin.com', '这是默认的超管账号', 1641440542, 0, 1);

# admin / 123456
