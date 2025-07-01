<?php

namespace App\Models\Admin;

use App\Models\Model;

class GroupUser extends Model
{
    protected $connection = 'default';
    protected $table = 'group_user';
    protected $fields = [
        'id' => 'int',
        'uid' => 'int',
        'gid' => 'int',
    ];

    public function getUids(array $gids): array
    {
        if (!$gids) {
            return [];
        }
        $list = $this->db()->select($this->table, ['uid'], ['gid' => $gids]);
        $list = $this->_settypes($list);
        return array_column($list, 'uid');
    }

    public function getGids(array $uids): array
    {
        if (!$uids) {
            return [];
        }
        $list = $this->db()->select($this->table, ['gid'], ['uid' => $uids]);
        $list = $this->_settypes($list);
        return array_column($list, 'gid');
    }

    public function setUserGroup(int $uid, array $gids): bool
    {
        if (!$uid) {
            return false;
        }
        $data = [];
        foreach ($gids as $v) {
            $data[] = ['uid' => $uid, 'gid' => $v];
        }
        // TODO: transaction
        // $this->db()->pdo->beginTransaction();
        // $this->db()->pdo->rollBack();
        // $this->db()->pdo->commit();
        $this->delete(['uid' => $uid]);
        $rowCount = $this->multiInsert($data);
        return !!$rowCount;
    }

    public function setGroupUser(int $gid, array $uids): bool
    {
        if (!$gid) {
            return false;
        }

        $data = [];
        foreach ($uids as $v) {
            $data[] = ['gid' => $gid, 'uid' => $v];
        }
        $this->delete(['gid' => $gid]);
        $rowCount = $this->multiInsert($data);
        return !!$rowCount;
    }
}
