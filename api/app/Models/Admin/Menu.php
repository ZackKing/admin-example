<?php

namespace App\Models\Admin;

use App\Models\Model;

use App\Components\Helper;

class Menu extends Model
{
    protected $connection = 'default';
    protected $table = 'menu';
    protected $fields = [
        'id' => 'int',
        'name' => 'string',
        'uri' => 'string',
        'level' => 'int',
        'pid' => 'int',
        'icon' => 'string',
        'status' => 'int',
        'sort' => 'int',
        'remark' => 'string',
        'created_time' => 'string',
        'updated_time' => 'string',
    ];

    const MAP_STSTUA = [
        'valid' => 1,
        'invalid' => 0,
    ];

    public function getAllMenu(bool $tree = true)
    {
        $list = $this->take();
        return $tree ? $this->_genTree($list) : $list;
    }

    public function getValidAllMenu(bool $tree = true)
    {
        $list = $this->take(['status' => self::MAP_STSTUA['valid']]);
        return $tree ? $this->_genTree($list) : $list;
    }

    public function getTree(array $mids): array
    {
        if (!$mids) {
            return [];
        }
        $list = $this->take(['id' => $mids, 'status' => self::MAP_STSTUA['valid']]);
        return $this->_genTree($list);
    }

    private function _genTree(array $list = []): array
    {
        if (!$list) {
            return [];
        }
        return Helper::genTree($list, 'id', 'pid', 'sub_menu');
    }
}
