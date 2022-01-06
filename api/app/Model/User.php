<?php

namespace App\Model;

class User extends Model
{
    protected $connection = 'default';
    protected $table = 'user';
    protected $fields = [
        'uid' => 'int',
        'name' => 'string',
        'password' => 'string',
        'salt' => 'string',
        'real_name' => 'string',
        'mobile' => 'string',
        'email' => 'string',
        'desc' => 'string',
        'login_time' => 'int',
        'pwd_wrong' => 'int',
        'status' => 'int',
        'created_time' => 'string',
        'updated_time' => 'string',
    ];

    const MAP_STSTUA = [
        'valid' => 1,
        'invalid' => 0,
    ];

    public function getUserMap(array $uids = []): array
    {
        if (!$uids) {
            return [];
        }
        $list = $this->take(['uid' => $uids]);
        return array_column($list, null, 'uid');
    }

    public function checkUserNameExist(string $userName, int $userId = 0): bool
    {
        return !!$this->count(['name' => $userName, 'uid[!]' => $userId]);
    }

}
