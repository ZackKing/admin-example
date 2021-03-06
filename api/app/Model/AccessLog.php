<?php

namespace App\Model;

class AccessLog extends Model
{
    protected $connection = 'default';
    protected $table = 'access_log';
    protected $fields = [
        'id' => 'int',
        'uid' => 'int',
        'method' => 'string',
        'path' => 'string',
        'header' => 'string',
        'query' => 'string',
        'body' => 'string',
        'ip' => 'string',
        'created_time' => 'string',
        'updated_time' => 'string',
    ];

}
