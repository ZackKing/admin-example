<?php

namespace App\Models\Admin;

use App\Models\Model;

class AccessLog extends Model
{
    protected $connection = 'default';
    protected $table = 'access_log';
    protected $fields = [
        'id' => 'int',
        'access_id' => 'string',
        'uid' => 'int',
        'method' => 'string',
        'path' => 'string',
        'header' => 'string',
        'query' => 'string',
        'body' => 'string',
        'ip' => 'string',
        'response' => 'string',
        'created_time' => 'string',
        'updated_time' => 'string',
    ];

}
