CREATE DATABASE `admin` DEFAULT CHARACTER SET = `utf8mb4`;

USE `admin`;

# user
CREATE TABLE `user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '',
    `password` varchar(255) NOT NULL DEFAULT '',
    `salt` varchar(255) NOT NULL DEFAULT '',
    `real_name` varchar(255) NOT NULL DEFAULT '',
    `mobile` varchar(255) NOT NULL DEFAULT '',
    `email` varchar(255) NOT NULL DEFAULT '',
    `desc` varchar(255) NOT NULL DEFAULT '',
    `login_time` int unsigned NOT NULL,
    `pwd_wrong` tinyint unsigned NOT NULL,
    `status` tinyint unsigned NOT NULL DEFAULT '1',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'Admin user';

INSERT INTO
    `user` (
        `uid`,
        `name`,
        `password`,
        `salt`,
        `real_name`,
        `mobile`,
        `email`,
        `desc`,
        `login_time`,
        `pwd_wrong`,
        `status`
    )
VALUES
    (
        1,
        'admin',
        '3a45ff59459512a21554bec8ea2e256b',
        '29e26F',
        'Super Admin',
        '18888888888',
        'admin@admin.com',
        'Default Super Admin',
        1641440542,
        0,
        1
    );

# group
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

INSERT INTO
    `group_user` (`id`, `uid`, `gid`)
VALUES
    (1, 1, 1);

CREATE TABLE `group` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL DEFAULT '',
    `status` tinyint(1) unsigned NOT NULL DEFAULT 1,
    `remark` varchar(200) NOT NULL DEFAULT '',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

# menu
CREATE TABLE `menu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '',
    `uri` varchar(100) DEFAULT '',
    `level` tinyint(2) unsigned DEFAULT '1',
    `pid` int(11) unsigned DEFAULT '0',
    `icon` varchar(100) DEFAULT '',
    `status` tinyint(1) unsigned DEFAULT '1',
    `sort` int(11) unsigned DEFAULT '999',
    `remark` varchar(255) DEFAULT '',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_pid` (`pid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO
    `menu` (
        `id`,
        `name`,
        `uri`,
        `level`,
        `pid`,
        `icon`,
        `status`,
        `sort`,
        `remark`
    )
VALUES
    (
        1,
        'Home',
        '/home',
        1,
        0,
        'dashboard',
        1,
        100,
        ''
    ),
    (
        2,
        'Dashboard',
        '/home/index',
        2,
        1,
        'dashboard',
        1,
        101,
        ''
    ),
    (
        3,
        'Management',
        '/panel',
        1,
        0,
        'config',
        1,
        900,
        ''
    ),
    (
        4,
        'User',
        '/panel/users',
        2,
        3,
        'user',
        1,
        901,
        ''
    ),
    (
        5,
        'Groups',
        '/panel/groups',
        2,
        3,
        'group',
        1,
        902,
        ''
    ),
    (
        6,
        'Menus',
        '/panel/menus',
        2,
        3,
        'tree',
        1,
        903,
        ''
    );

# log
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