<?php

namespace App\Http\Controllers;

use App\Components\Validator;
use App\Exceptions\LogicException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Laravel\Lumen\Routing\Controller as BaseController;

class Controller extends BaseController
{

    public function ok($data = true): JsonResponse
    {
        return response()->json([
            'code' => 0,
            'data' => $data,
            'msg' => 'success',
            'ts' => time(),
        ]);
    }

    public function fail(int $code = -1, string $msg = '', array $data = [], array $params = [])
    {
        throw new LogicException($code, $msg, $data, $params);
    }

    public function getUid(Request $r)
    {
        return $r->offsetGet('sys_uid', 0);
    }

    /**
     * get from query prarms
     * @param Request $r
     * @param array $fieldList
     * @return array
     */
    public function getQuery(Request $r, array $fieldList = []): array
    {
        return Validator::valid($r->query->all(), $fieldList);
    }

    /**
     * get from json body
     * @param Request $r
     * @param array $fieldList
     * @return array
     */
    public function getData(Request $r, array $fieldList = []): array
    {
        return Validator::valid($r->json()->all(), $fieldList);
    }

    /**
     * get from formdata
     * @param Request $r
     * @param array $fieldList
     * @return array
     */
    public function getForm(Request $r, array $fieldList = []): array
    {
        return validator::valid($r->input(), $fieldList);
    }

    /**
     * 校验数据
     * @param array $data
     * @param array $fieldList
     * @return array
     */
    protected function validData(array $data, array $fieldList = []): array
    {
        return Validator::valid($data, $fieldList);
    }

}
