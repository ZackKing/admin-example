<?php

namespace App\Model;

class GroupMenu extends Model
{
    protected $connection = 'default';
    protected $table = 'group_menu';
    protected $fields = [
        'mid' => 'int',
        'gid' => 'int',
    ];

    public function getMids(array $gids): array
    {
        if (!$gids) {
            return [];
        }
        $list = $this->take(['gid' => $gids], ['mid']);
        $list = $this->_settypes($list);
        return array_column($list, 'mid');
    }

    public function getGids(array $mids): array
    {
        if (!$mids) {
            return [];
        }
        $list = $this->take(['mid' => $mids], ['gid']);
        $list = $this->_settypes($list);
        return array_column($list, 'gid');
    }

    public function setMenuGroup(int $mid, array $gids): bool
    {
        if (!$mid) {
            return false;
        }
        $data = [];
        foreach ($gids as $v) {
            $data[] = ['mid' => $mid, 'gid' => $v];
        }
        // TODO: transaction
        // $this->db()->pdo->beginTransaction();
        // $this->db()->pdo->rollBack();
        // $this->db()->pdo->commit();
        $this->delete(['mid' => $mid]);
        $rowCount = $this->multiInsert($data);
        return !!$rowCount;
    }

    public function setGroupMenu(int $gid, array $mids = []): bool
    {
        if (!$gid) {
            return false;
        }

        $data = [];
        foreach ($mids as $v) {
            $data[] = ['gid' => $gid, 'mid' => $v];
        }
        $this->delete(['gid' => $gid]);
        $rowCount = $this->multiInsert($data);
        return !!$rowCount;
    }

}
