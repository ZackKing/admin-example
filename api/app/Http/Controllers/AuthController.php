<?php

namespace App\Http\Controllers;

use App\Components\Helper;
use App\Http\Logic\Auth as AuthLogic;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Validation\Rule;

class AuthController extends Controller
{

    /**
     * @api {post} /login Login - login
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "account": "account",
     *      "password": "password",
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": {
     *          "token": "jwt_token_string"
     *      },
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function login(Request $r): JsonResponse
    {
        $data = $this->getData($r, [
            'account' => ['rule' => 'required|string'],
            'password' => ['rule' => 'required|string'],
        ]);

        $authLogic = AuthLogic::getInstance();
        $ip = Helper::realIp();

        $token = $authLogic->login($data['account'], $data['password']);
        $authLogic->loginOk($data['account'], $ip);
        return $this->retSuccess([
            'token' => $token,
        ]);
    }

    /**
     * @api {get} /renewToken Login - renew jwt token
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": {
     *          "token": "jwt_token_string"
     *      },
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function renewToken(Request $request): JsonResponse
    {
        return $this->retSuccess([
            'token' => AuthLogic::getInstance()->createToken($request->jwt['uid']),
        ]);
    }

    /**
     * @api {get} /auth/menu Permission - tree of menu
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": [
     *          {
     *               "id": 1,
     *               "name": "PP Assistant",
     *               "uri": "",
     *               "icon": "",
     *               "sort": 10,
     *               "sub_menu": [] // leaf node is empty
     *           }
     *      ],
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function menu(Request $request): JsonResponse
    {
        $menuTree = AuthLogic::getInstance()->getUserMenu($this->getUid($request));
        return $this->retSuccess($menuTree);
    }

    /**
     * @api {post} /auth/password Password - change
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParamExample {json} Request-Example:
     *  {
     *      "password": "new_pwd",
     *      "old_password": "old_pwd",
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": true,
     *    "msg": "success"
     *  }
     */
    public function changPwd(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'password' => ['rule' => 'required|string|min:6'],
            'old_password' => ['rule' => 'required|string'],
        ]);
        $uid = $this->getUid($request);
        $res = AuthLogic::getInstance()->changePwd($uid, $data['password'], $data['old_password']);
        if (!$res) {
            $this->retError(10004);
        }
        return $this->retSuccess();
    }

    /**
     * @api {post} /group/user Group - set user
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "id": 1, // required, group id
     *      "uids": [1,2,3] // user id array
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": true,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function setGroupUser(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'id' => ['rule' => 'required|int|min:0'],
            'uids' => ['rule' => 'array', 'default' => []],
        ]);
        AuthLogic::getInstance()->setGroupUser($data['id'], $data['uids']);
        return $this->retSuccess(true);
    }

    /**
     * @api {post} /group/menu Group - set menu
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "id": 1, // group id
     *      "menu_ids": [1,2,3] // menu id array
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": true,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function setGroupMenu(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'id' => ['rule' => 'required|int|min:0'],
            'menu_ids' => ['rule' => 'array', 'default' => []],
        ]);
        AuthLogic::getInstance()->setGroupMenu($data['id'], $data['menu_ids']);
        return $this->retSuccess(true);
    }

    /**
     * @api {post} /menu/group Menu - set group
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "id": 1, // required, menu id
     *      "group_ids": [1,2,3] // group id array
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": true,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function setMenuGroup(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'id' => ['rule' => 'required|int|min:0'],
            'group_ids' => ['rule' => 'array', 'default' => []],
        ]);
        AuthLogic::getInstance()->setMenuGroup($data['id'], $data['group_ids']);
        return $this->retSuccess(true);
    }

    /**
     * @api {get} /menu Menu - tree of all menu
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": [
     *          {
     *               "menu_id": 1,
     *               "menu_name": "PP Assistant",
     *               "uri": "",
     *               "icon": "",
     *               "sort": 10,
     *               "sub_menu": [] // leaf node is empty
     *           }
     *      ],
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function menuTree(Request $request): JsonResponse
    {
        $tree = AuthLogic::getInstance()->menuTree();
        return $this->retSuccess($tree);
    }

    /**
     * @api {get} /menu/info Menu - info
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParam {Number} id menu id
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": {
     *          "id": 1,
     *          "name": "PP Assistant",
     *          "uri": "",
     *          "icon": "",
     *          "sort": 10,
     *          "group_ids": []
     *      },
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function menuInfo(Request $request): JsonResponse
    {
        $query = $this->getParam($request, [
            'id' => ['rule' => 'required|int|min:0'],
        ]);
        $info = AuthLogic::getInstance()->menuInfo($query['id']);
        return $this->retSuccess($info);
    }

    /**
     * @api {get} /group Group - list
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": [
     *        {
     *            "id": 1,
     *            "name": "admin",
     *            "status": 1,
     *            "remark": "super admin have all access",
     *            "uids": [],
     *            "mids": [],
     *            "created_time": "2020-06-02 15:20:40",
     *            "updated_time": "2020-06-02 15:21:39"
     *        }
     *    ],
     *    "msg": "success"
     *}
     */
    public function groupList(Request $request): JsonResponse
    {
        $list = AuthLogic::getInstance()->groupList();
        return $this->retSuccess($list);
    }

    /**
     * @api {post} /group Group - add
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "name": "group name",
     *      "remark": "group remark"
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": 1,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function addGroup(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'name' => ['rule' => 'required|string|min:2'],
            'remark' => ['rule' => 'string', 'default' => ''],
        ]);
        $id = AuthLogic::getInstance()->setGroup(0, $data);
        return $this->retSuccess($id);
    }
    /**
     * @api {post} /group/edit Group - edit
     * @apiGroup Auth
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "id": 123, // required
     *      "name": "group name",
     *      "remark": "group remark",
     *      "status": "group status"
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": true,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function editGroup(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'id' => ['rule' => 'required|int|min:0'],
            'name' => ['rule' => 'string|min:2'],
            'remark' => ['rule' => 'string'],
            'status' => ['rule' => [Rule::in([0, 1])]],
        ]);
        $info = Helper::fieldFilter($data, ['name', 'remark', 'status']);
        if ($info) {
            AuthLogic::getInstance()->setGroup($data['id'], $info);
        }
        return $this->retSuccess(true);
    }

}
