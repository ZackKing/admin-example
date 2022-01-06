<?php

return [
    'default' => [
        'database_type' => 'mysql',
        'write' => [
            'database_name' => 'admin',
            'server' => 'mysql',
            'username' => 'root',
            'password' => 'admin123456',
            'port' => 3306,
        ],
        'read' => [
            'database_name' => 'admin',
            'server' => 'mysql',
            'username' => 'root',
            'password' => 'admin123456',
            'port' => 3306,
        ],
        // ----------------------- [optional] -----------------------
        'charset' => 'utf8',
        'collation' => 'utf8_general_ci',
        'prefix' => '',
        'logging' => false,
        'option' => [ // driver_option for connection, read more from http://www.php.net/manual/en/pdo.setattribute.php
            PDO::ATTR_CASE => PDO::CASE_NATURAL,
            PDO::MYSQL_ATTR_INIT_COMMAND => 'SET @@session.time_zone= \'+08:00\'',
        ],
        'command' => [ // Medoo will execute those commands after connected to the database for initialization
            // 'SET SQL_MODE=ANSI_QUOTES',
        ],
        // ----------------------- [optional] -----------------------
    ],
];
