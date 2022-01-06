<?php

namespace App\Model;

class Group extends Model
{
    protected $connection = 'default';
    protected $table = 'group';
    protected $fields = [
        'id' => 'int',
        'name' => 'string',
        'status' => 'int',
        'remark' => 'string',
        'created_time' => 'string',
        'updated_time' => 'string',
    ];

    const ADMIN_ID = 1;

    public function isAdmin(array $gids): bool
    {
        return \in_array(self::ADMIN_ID, $gids);
    }

    protected function _parseList(array $list): array
    {
        foreach ($list as $k => $v) {
            $list[$k] = $this->_parse($v);
        }
        return $list;
    }

    protected function _parse(array $data): array
    {
        return $data;
    }

}
