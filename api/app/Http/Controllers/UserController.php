<?php

namespace App\Http\Controllers;

use App\Components\Helper;
use App\Http\Logic\Auth as AuthLogic;
use App\Http\Logic\User as UserLogic;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Validation\Rule;

class UserController extends Controller
{
    /**
     * @api {get} /user/self Self - info
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": {
     *        "uid": 1,
     *        "name": "admin",
     *        "real_name": "SomewhereYu",
     *        "mobile": "13800138001",
     *        "email": "admin@osadmin.org",
     *        "desc": "初始的超级管理员!!",
     *        "login_time": 1591068875,
     *        "status": 1,
     *        "login_ip": "211.97.6.100",
     *        "user_group": 1,
     *        "template": "default",
     *        "shortcuts": "1,7,10,11,13,14,18,21,22,24,111,137,142,143,152,188",
     *        "show_quicknote": 0,
     *        "code": "0"
     *    },
     *    "msg": "success"
     *  }
     */
    public function self(Request $r)
    {
        $info = UserLogic::instance()->info($this->getUid($r));
        if (isset($info['password'])) {
            unset($info['password']);
        }
        return $this->ok($info);
    }

    /**
     * @api {post} /user/self Self - edit info
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParamExample {json} Request-Example:
     *  {
     *      "real_name": "real_name",
     *      "mobile": "mobile",
     *      "email": "email@email.com",
     *      "desc": "descdescdescdesc"
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": true,
     *    "msg": "success"
     *  }
     */
    public function editSelf(Request $request): JsonResponse
    {
        $uid = $this->getUid($request);
        $data = $this->getData($request, [
            'real_name' => ['rule' => 'string'],
            'mobile' => ['rule' => 'string'],
            'email' => ['rule' => 'string'],
            'desc' => ['rule' => 'string'],
        ]);
        if ($data) {
            UserLogic::instance()->edit($uid, $data);
        }
        return $this->ok();
    }

    /**
     * @api {get} /user/list Account - list
     * @apiGroup User
     * @apiDescription author: zack
     * @apiUse BaseApiDefined
     * @apiParam {integer} [size=10] response data coun
     * @apiParam {integer} [page=1] page number
     * @apiParam {String} [name] login user name
     * @apiParam {integer[]} [group_ids] group id array
     * @apiSuccessExample {json} Success-Response:
     *    {
     *    "code": 0,
     *    "data": {
     *        "list": [
     *            {
     *                "uid": 2,
     *                "name": "test",
     *                "real_name": "dachui",
     *                "desc": "dev test",
     *                "status": 1,
     *            }
     *        ],
     *        "total": 1,
     *        "size": 10,
     *        "page": 1
     *    },
     *    "msg": "success"
     *}
     */
    function list(Request $request): JsonResponse {
        $query = $this->getQuery($request, [
            'name' => ['rule' => 'string', 'ignore' => ['']],
            'group_ids' => ['rule' => 'array'],
            'size' => ['rule' => 'int', 'default' => 10],
            'page' => ['rule' => 'int', 'default' => 1],
        ]);
        $query['page'] = (int) $query['page'];
        $query['size'] = (int) $query['size'];

        $option = [
            'limit' => $query['size'],
            'offset' => $query['size'] * ($query['page'] - 1),
            'columns' => ['uid', 'name', 'real_name', 'mobile', 'email', 'desc', 'status'],
        ];
        $search = [
            'status' => [0, 1],
        ];
        !empty($query['name']) && $search['name'] = $query['name'];
        !empty($query['group_ids']) && $search['group_ids'] = $query['group_ids'];

        $data = UserLogic::instance()->search($search, $option);
        if ($data['list']) {
            $data['list'] = UserLogic::instance()->addGroupInfo($data['list'], ['id', 'name', 'status']);
        }
        $data['size'] = $query['size'];
        $data['page'] = $query['page'];
        return $this->ok($data);
    }

    /**
     * @api {get} /user/info Account - info
     * @apiGroup User
     * @apiDescription author: zack
     * @apiUse BaseApiDefined
     * @apiParam {integer} uid account user id
     * @apiSuccessExample {json} Success-Response:
     *    {
     *    "code": 0,
     *    "data": {
     *        "uid": 2,
     *        "name": "test",
     *        "real_name": "dachui",
     *        "desc": "dev test",
     *        "status": 1,
     *    },
     *    "msg": "success"
     *}
     */
    public function info(Request $request): JsonResponse
    {
        $query = $this->getQuery($request, [
            'uid' => ['rule' => 'required|int'],
        ]);
        $columns = ['uid', 'name', 'real_name', 'mobile', 'email', 'desc', 'status'];
        $info = UserLogic::instance()->info($query['uid'], $columns);
        if ($info) {
            $list = UserLogic::instance()->addGroupInfo([$info], ['id', 'name', 'status']);
            $info = $list[0];
        }
        return $this->ok($info);
    }

    /**
     * @api {post} /user/update Account - edit
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParamExample {json} Request-Example:
     *  {
     *      "uid": 123, // required
     *      "real_name": "real_name",
     *      "mobile": "mobile",
     *      "email": "email@email.com",
     *      "desc": "descdescdescdesc",
     *      "password": ""
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": true,
     *    "msg": "success"
     *  }
     */
    public function updateInfo(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'uid' => ['rule' => 'required|int'],
            'name' => ['rule' => 'string', 'ignore' => ['']],
            'password' => ['rule' => 'string', 'ignore' => ['']],
            'real_name' => ['rule' => 'string'],
            'mobile' => ['rule' => 'string'],
            'email' => ['rule' => 'string'],
            'desc' => ['rule' => 'string'],
            'department' => ['rule' => 'string'],
        ]);

        $info = Helper::fieldFilter($data, ['name', 'real_name', 'mobile', 'email', 'desc', 'department']);
        if ($info) {
            UserLogic::instance()->edit($data['uid'], $info);
        }
        if (!empty($data['password'])) {
            AuthLogic::instance()->setPwd($data['uid'], $data['password']);
        }
        return $this->ok(true);
    }

    /**
     * @api {post} /user/status Account - set status
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParamExample {json} Request-Example:
     *  {
     *      "uid": 123, // required
     *      "status": 0,  // 0 - ban,  1 - normal
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": true,
     *    "msg": "success"
     *  }
     */
    public function status(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'uid' => ['rule' => 'required|int'],
            'status' => ['rule' => ['required', Rule::in([0, 1, 2])]],
        ]);

        UserLogic::instance()->edit($data['uid'], ['status' => $data['status']]);
        return $this->ok(true);
    }

    /**
     * @api {post} /user/add Account - add
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiUse BaseApiDefined
     * @apiParamExample {json} Request-Example:
     *  {
     *      "name": "login name", // required, min: 6, max: 32
     *      "password": "", // required, min: 6
     *      "real_name": "real_name", // required
     *      "mobile": "mobile", // required
     *      "email": "email@email.com", // required
     *      "desc": "descdescdescdesc"
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *    "code": 0,
     *    "data": true,
     *    "msg": "success"
     *  }
     */
    public function add(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'name' => ['rule' => 'required|string|min:6'],
            'password' => ['rule' => 'required|string|min:6'],
            // 'real_name' => ['rule' => 'required|string'],
            // 'mobile' => ['rule' => 'required|string'],
            'email' => ['rule' => 'required|string'],
            'desc' => ['rule' => 'string'],
        ]);
        $uid = UserLogic::instance()->add($data);
        if (!$uid) {
            $this->fail(10006);
        }
        return $this->ok(['uid' => $uid]);
    }

    /**
     * @api {post} /user/group Account - set group
     * @apiGroup User
     * @apiDescription author: Zack
     * @apiParamExample {json} Request-Example:
     *  {
     *      "id": 1, // required, group id
     *      "group_ids": [1,2,3] // group id array
     *  }
     * @apiSuccessExample {json} Success-Response:
     *  {
     *      "data": true,
     *      "code": 0,
     *      "msg": "success"
     *   }
     */
    public function setGroup(Request $request): JsonResponse
    {
        $data = $this->getData($request, [
            'id' => ['rule' => 'required|int|min:0'],
            'group_ids' => ['rule' => 'array', 'default' => []],
        ]);
        AuthLogic::instance()->setUserGroup($data['id'], $data['group_ids']);
        return $this->ok(true);
    }

}
