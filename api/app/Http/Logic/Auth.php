<?php

namespace App\Http\Logic;

use App\Components\Helper;
use App\Components\JWTHelper;
use App\Models\Admin\Group;
use App\Models\Admin\GroupMenu;
use App\Models\Admin\GroupUser;
use App\Models\Admin\Menu;
use App\Models\Admin\User;

class Auth extends Logic
{
    const PWD_WRONG_LIMIT = 5;

    public function login(string $account, string $pwd): string
    {
        $user = $this->getAccountUser($account);
        if (empty($user['password']) || $user['password'] !== $this->encodePwd($pwd, $user['salt'])) {
            $limit = $this->checkBan($user);
            $this->throw(10008, '', [], [$limit]);
        }
        return JWTHelper::createToken($user['uid']);
    }

    public function getUserMenu(int $uid): array
    {
        if ($uid < 1) {
            return [];
        }
        $menuTree = [];
        $gids = GroupUser::instance()->getGids([$uid]);
        if ($this->_isAdmin($gids)) { // super admin
            $menuTree = Menu::instance()->getValidAllMenu();
        } else {
            $menuIds = GroupMenu::instance()->getMids($gids);
            $menuTree = Menu::instance()->getTree($menuIds);
        }

        return $menuTree;
    }

    public function changePwd(int $uid, string $pwd, string $old_pwd): bool
    {
        if (!$uid || !$pwd || !$old_pwd) {
            return false;
        }
        $user = User::instance()->takeFirst(['uid' => $uid]);
        if (empty($user['password']) || $user['password'] !== $this->encodePwd($old_pwd, $user['salt'])) {
            return false;
        }
        $salt = $this->genSalt();
        return !!User::instance()->update(['uid' => $uid], ['password' => $this->encodePwd($pwd, $salt), 'salt' => $salt]);
    }

    public function setPwd(int $uid, string $pwd): bool
    {
        if (!$uid || !$pwd) {
            return false;
        }
        $salt = $this->genSalt();
        return !!User::instance()->update(['uid' => $uid], ['password' => $this->encodePwd($pwd, $salt), 'salt' => $salt]);
    }

    public function encodePwd(string $pwd, string $salt): string
    {
        return md5($salt . md5($pwd));
    }

    public function setUserGroup(int $uid, array $gids): bool
    {
        return GroupUser::instance()->setUserGroup($uid, $gids);
    }

    public function groupList()
    {
        $list = Group::instance()->take();
        // TODO: need optimization
        foreach ($list as $k => $v) {
            $list[$k]['uids'] = GroupUser::instance()->getUids([$v['id']]);
            $list[$k]['mids'] = GroupMenu::instance()->getMids([$v['id']]);
        }
        return $list;
    }

    public function setGroup(int $id, array $data): int
    {
        if ($id > 0) { // update
            Group::instance()->update(['id' => $id], $data);
        } else if ($id == 0) { // insert
            $id = Group::instance()->insert($data);
        }
        return $id;
    }

    public function setGroupUser(int $gid, array $uids): bool
    {
        return GroupUser::instance()->setGroupUser($gid, $uids);
    }

    public function setGroupMenu(int $gid, array $mids): bool
    {
        return GroupMenu::instance()->setGroupMenu($gid, $mids);
    }

    public function menuTree(): array
    {
        return Menu::instance()->getAllMenu();
    }

    public function menuInfo(int $id): array
    {
        $info = Menu::instance()->takeFirst(['id' => $id]);
        if ($info) {
            $info['group_ids'] = GroupMenu::instance()->getGids([$id]);
        }
        return $info;
    }

    public function setMenuGroup(int $mid, array $gids): bool
    {
        return GroupMenu::instance()->setMenuGroup($mid, $gids);
    }

    public function checkBan(array $user): int
    {
        ++$user['pwd_wrong'];
        if ($user['pwd_wrong'] >= self::PWD_WRONG_LIMIT) { // set status ban
            User::instance()->update(
                ['uid' => $user['uid']],
                ['status' => User::MAP_STSTUA['invalid'], 'pwd_wrong' => $user['pwd_wrong']]
            );
            $this->throw(10009);
        } else {
            User::instance()->update(['uid' => $user['uid']], ['pwd_wrong' => $user['pwd_wrong']]);
        }
        return self::PWD_WRONG_LIMIT - $user['pwd_wrong'];
    }

    public function loginOk(string $account, string $ip)
    {
        User::instance()->update(['name' => $account], ['login_ip' => $ip, 'pwd_wrong' => 0, 'login_time' => time()]);
    }

    public function getAccountUser(string $account): array
    {
        $user = User::instance()->takeFirst(['name' => $account]);
        if (empty($user)) {
            $this->throw(10003);
        }
        if (!$user['status']) {
            $this->throw(10005);
        }
        return $user;
    }

    public function genSalt(): string
    {
        return Helper::randStr(6);
    }

    private function _isAdmin(array $gids): bool
    {
        return Group::instance()->isAdmin($gids);
    }

}
