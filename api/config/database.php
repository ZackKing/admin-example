<?php

return [
    'db' => [
        'default' => [
            'type' => env('DB_CONNECTION', 'mysql'),
            'host' => env('DB_HOST', '127.0.0.1'),
            'database' => env('DB_DATABASE', 'admin'),
            'username' => env('DB_USERNAME', ''),
            'password' => env('DB_PASSWORD', ''),
            'port' => env('DB_PORT', 3306),
            'charset' => env('DB_CHARSET', 'utf8mb4'),
            'collation' => env('DB_COLLATION', 'utf8mb4_general_ci'),
            'prefix' => env('DB_PREFIX', ''),
            'logging' => env('APP_DEBUG', false),
            'option' => [
                PDO::ATTR_CASE => PDO::CASE_NATURAL,
                // PDO::MYSQL_ATTR_INIT_COMMAND => 'SET @@session.time_zone= \'+08:00\'',
            ],
            'command' => [],
        ],
        // 'default' => [ // 读写分离配置方法
        //     'type' => env('DB_CONNECTION', 'mysql'),
        //     'write' => [ // 写
        //         'host' => env('DB_HOST', '127.0.0.1'),
        //         'database' => env('DB_DATABASE', 'admin'),
        //         'username' => env('DB_USERNAME', ''),
        //         'password' => env('DB_PASSWORD', ''),
        //         'port' => env('DB_PORT', 3306),
        //     ],
        //     'read' => [ // 读
        //         'host' => env('DB_HOST_READ', env('DB_HOST', '127.0.0.1')),
        //         'database' => env('DB_DATABASE', 'admin'),
        //         'username' => env('DB_USERNAME_READ', env('DB_USERNAME', '')),
        //         'password' => env('DB_PASSWORD_READ', env('DB_PASSWORD', '')),
        //         'port' => env('DB_PORT_READ', env('DB_PORT', 3306)),
        //     ],
        //     'charset' => env('DB_CHARSET', 'utf8mb4'),
        //     'collation' => env('DB_COLLATION', 'utf8mb4_general_ci'),
        //     'prefix' => env('DB_PREFIX', ''),
        //     'logging' => false,
        //     'option' => [
        //         PDO::ATTR_CASE => PDO::CASE_NATURAL,
        //         // PDO::MYSQL_ATTR_INIT_COMMAND => 'SET @@session.time_zone= \'+08:00\'',
        //     ],
        //     'command' => [],
        // ],
    ],

    'redis' => [
        'default' => [
            'host' => env('REDIS_HOST', '127.0.0.1'),
            'port' => env('REDIS_PORT', 6379),
            'db' => env('REDIS_DB', 0),
            'password' => env('REDIS_PASSWORD', null),
        ],

        'cache' => [
            'host' => env('REDIS_HOST', '127.0.0.1'),
            'port' => env('REDIS_PORT', 6379),
            'db' => env('REDIS_CACHE_DB', 1),
            'password' => env('REDIS_PASSWORD', null),
        ],
    ],

];
