<?php

return [
    'name' => env('APP_NAME', 'admin'),
    'env' => env('APP_ENV', 'dev'),
    'key' => env('APP_KEY', 'test_key'),
    'debug' => env('APP_DEBUG', false),
    'url' => env('APP_URL', 'xxxx.com'),
    'timezone' => env('APP_TIMEZONE', 'Asia/Hong_Kong'),
    'jwt' => [
        'exp' => env('AUTH_JWT_EXP', 604800),
        'iss' => env('AUTH_JWT_ISS', env('APP_URL', 'xxx.com')),
        'secret' => env('AUTH_JWT_SECRET', 'secret_for_example'),
        'algorithm' => env('AUTH_JWT_ALGORITHM', 'HS256'),
    ],
];
