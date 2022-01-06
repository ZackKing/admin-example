<?php

namespace App\Http\Logic;

use App\Components\Helper;
use App\Model\Group;
use App\Model\GroupUser;
use App\Model\User as UserMdl;

class User extends Logic
{
    public function info(int $uid, array $cloumns = []): array
    {
        return UserMdl::getInstance()->takeFirst(['uid' => $uid], $cloumns);
    }

    public function edit(int $uid, array $data): int
    {
        if (!empty($data['name'])) { // check name unique
            if (UserMdl::getInstance()->checkUserNameExist($data['name'], $uid)) {
                $this->throwHttpError(10007);
            }
        }
        return UserMdl::getInstance()->update(['uid' => $uid], $data);
    }

    public function search(array $search = [], array $option = []): array
    {
        empty($option['limit']) && $option['limit'] = 10;
        empty($option['offset']) && $option['offset'] = 0;
        empty($option['orderBy']) && $option['orderBy'] = ['uid' => 'DESC'];
        $data = ['list' => [], 'total' => 0];
        $where = [];
        foreach ($search as $k => $v) {
            switch ($k) {
                case 'group_ids':
                    $uids = GroupUser::getInstance()->getUids($search['group_ids']);
                    if (!$uids) {
                        return $data;
                    }
                    $where['uid'] = $uids;
                    break;
                default:
                    $where[$k] = $v;
                    break;
            }
        }
        $mdl = UserMdl::getInstance();
        $data['total'] = $mdl->count($where);
        if ($data['total'] > 0) {
            $where['LIMIT'] = [$option['offset'], $option['limit']];
            $where['ORDER'] = $option['orderBy'];
            empty($option['columns']) && $option['columns'] = $mdl->getColumns();
            $data['list'] = UserMdl::getInstance()->take(
                $where,
                $option['columns']
            );
        }
        return $data;
    }

    public function addGroupInfo(array $list, array $cloumns = []): array
    {
        $uids = array_column($list, 'uid');
        if (!$uids) {
            return [];
        }
        $groupUser = GroupUser::getInstance()->take(['uid' => $uids]);
        $gids = array_column($groupUser, 'gid');
        $groupList = Group::getInstance()->take(['id' => $gids], $cloumns);
        $groupHash = Helper::value2key($groupList, 'id');
        $hash = Helper::value2key($list, 'uid');
        foreach ($groupUser as $k => $v) {
            empty($hash[$v['uid']]['group']) && $hash[$v['uid']]['group'] = [];
            if (!empty($groupHash[$v['gid']])) {
                $hash[$v['uid']]['group'][] = $groupHash[$v['gid']];
            }
        }
        return Helper::hash2list($hash);
    }

    public function add(array $data): int
    {
        $mdl = UserMdl::getInstance();
        if ($mdl->checkUserNameExist($data['name'], 0)) {
            $this->throwHttpError(10007);
        }
        $authLogic = Auth::getInstance();
        $data['salt'] = $authLogic->genSalt();
        $data['password'] = $authLogic->encodePwd($data['password'], $data['salt']);
        $data['status'] = UserMdl::MAP_STSTUA['valid'];
        return $mdl->insert($data);
    }

}
