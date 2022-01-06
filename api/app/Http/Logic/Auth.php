<?php

namespace App\Http\Logic;

use App\Components\Helper;
use App\Model\Group;
use App\Model\GroupMenu;
use App\Model\GroupUser;
use App\Model\Menu;
use App\Model\User;
use Firebase\JWT\JWT;

class Auth extends Logic
{
    const PWD_WRONG_LIMIT = 5;

    public function login(string $account, string $pwd): string
    {
        $user = $this->getAccountUser($account);
        if (empty($user['password']) || $user['password'] !== $this->encodePwd($pwd, $user['salt'])) {
            $limit = $this->checkBan($user);
            $this->throwHttpError(10008, '', [], [$limit]);
        }
        return $this->createToken($user['uid']);
    }

    public function createToken(int $uid = 0): string
    {
        if (!$uid) {
            return '';
        }
        $conf = config('jwt');
        $signTime = time();
        $token = JWT::encode([
            'uid' => $uid,
            'iss' => $conf['iss'],
            'iat' => $signTime,
            'exp' => $signTime + $conf['exp'],
        ],
            $conf['secretKey'],
            $conf['algorithms']
        );
        return $token;
    }

    public function parseToken(string $token)
    {
        $conf = config('jwt');
        try {
            $decoded = JWT::decode($token, $conf['secretKey'], [$conf['algorithms']]);
            return Helper::obj2arr($decoded);
        } catch (\Throwable$th) {
            $this->throwHttpError(10002);
            return [];
        }
    }

    public function getUserMenu(int $uid): array
    {
        if ($uid < 1) {
            return [];
        }
        $menuTree = [];
        $gids = GroupUser::getInstance()->getGids([$uid]);
        if ($this->_isAdmin($gids)) { // super admin
            $menuTree = Menu::getInstance()->getValidAllMenu();
        } else {
            $menuIds = GroupMenu::getInstance()->getMids($gids);
            $menuTree = Menu::getInstance()->getTree($menuIds);
        }

        return $menuTree;
    }

    public function changePwd(int $uid, string $pwd, string $old_pwd): bool
    {
        if (!$uid || !$pwd || !$old_pwd) {
            return false;
        }
        $user = User::getInstance()->takeFirst(['uid' => $uid]);
        if (empty($user['password']) || $user['password'] !== $this->encodePwd($old_pwd, $user['salt'])) {
            return false;
        }
        $salt = $this->genSalt();
        return !!User::getInstance()->update(['uid' => $uid], ['password' => $this->encodePwd($pwd, $salt), 'salt' => $salt]);
    }

    public function setPwd(int $uid, string $pwd): bool
    {
        if (!$uid || !$pwd) {
            return false;
        }
        $salt = $this->genSalt();
        return !!User::getInstance()->update(['uid' => $uid], ['password' => $this->encodePwd($pwd, $salt), 'salt' => $salt]);
    }

    public function encodePwd(string $pwd, string $salt): string
    {
        return md5($salt . md5($pwd));
    }

    public function setUserGroup(int $uid, array $gids): bool
    {
        return GroupUser::getInstance()->setUserGroup($uid, $gids);
    }

    public function groupList()
    {
        $list = Group::getInstance()->take();
        // TODO: need optimization
        foreach ($list as $k => $v) {
            $list[$k]['uids'] = GroupUser::getInstance()->getUids([$v['id']]);
            $list[$k]['mids'] = GroupMenu::getInstance()->getMids([$v['id']]);
        }
        return $list;
    }

    public function setGroup(int $id, array $data): int
    {
        if ($id > 0) { // update
            Group::getInstance()->update(['id' => $id], $data);
        } else if ($id == 0) { // insert
            $id = Group::getInstance()->insert($data);
        }
        return $id;
    }

    public function setGroupUser(int $gid, array $uids): bool
    {
        return GroupUser::getInstance()->setGroupUser($gid, $uids);
    }

    public function setGroupMenu(int $gid, array $mids): bool
    {
        return GroupMenu::getInstance()->setGroupMenu($gid, $mids);
    }

    public function menuTree(): array
    {
        return Menu::getInstance()->getAllMenu();
    }

    public function menuInfo(int $id): array
    {
        $info = Menu::getInstance()->takeFirst(['id' => $id]);
        if ($info) {
            $info['group_ids'] = GroupMenu::getInstance()->getGids([$id]);
        }
        return $info;
    }

    public function setMenuGroup(int $mid, array $gids): bool
    {
        return GroupMenu::getInstance()->setMenuGroup($mid, $gids);
    }

    public function checkBan(array $user): int
    {
        ++$user['pwd_wrong'];
        if ($user['pwd_wrong'] >= self::PWD_WRONG_LIMIT) { // set status ban
            User::getInstance()->update(
                ['uid' => $user['uid']],
                ['status' => User::MAP_STSTUA['invalid'], 'pwd_wrong' => $user['pwd_wrong']]
            );
            $this->throwHttpError(10009);
        } else {
            User::getInstance()->update(['uid' => $user['uid']], ['pwd_wrong' => $user['pwd_wrong']]);
        }
        return self::PWD_WRONG_LIMIT - $user['pwd_wrong'];
    }

    public function loginOk(string $account, string $ip)
    {
        User::getInstance()->update(['name' => $account], ['login_ip' => $ip, 'pwd_wrong' => 0, 'login_time' => time()]);
    }

    public function getAccountUser(string $account): array
    {
        $user = User::getInstance()->takeFirst(['name' => $account]);
        if (empty($user)) {
            $this->throwHttpError(10003);
        }
        if (!$user['status']) {
            $this->throwHttpError(10005);
        }
        return $user;
    }

    public function genSalt(): string
    {
        return Helper::randStr(6);
    }

    private function _isAdmin(array $gids): bool
    {
        return Group::getInstance()->isAdmin($gids);
    }

}
