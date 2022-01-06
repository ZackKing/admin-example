<?php

namespace App\Http\Controllers;

use App\Components\HttpError;
use App\Components\Validator;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Laravel\Lumen\Routing\Controller as BaseController;

class Controller extends BaseController
{

    public function retSuccess($data = true): JsonResponse
    {
        return response()->json([
            'code' => 0,
            'data' => $data,
            'msg' => 'success',
        ]);
    }

    public function retError(int $code = -1, string $msg = '', array $data = [], array $params = [])
    {
        throw new HttpError($code, $msg, $data, $params);
    }

    public function getParam(Request $request, array $fieldList = []): array
    {
        return $this->_validData($request->query->all(), $fieldList);
    }

    public function getData(Request $request, array $fieldList = []): array
    {
        return $this->_validData($request->json()->all(), $fieldList);
    }

    protected function getUid(Request $request): int
    {
        return empty($request->jwt['uid']) ? 0 : (int) $request->jwt['uid'];
    }

    protected function getRealIP(Request $r): string
    {
        return $r->realIp ?? '';
    }

    protected function validData(array $data, array $fieldList = []): array
    {
        return $this->_validData($data, $fieldList);
    }

    private function _validData(array $data, array $fieldList = []): array
    {
        $rules = [];
        $safeData = [];
        foreach ($fieldList as $k => $v) {
            $rules[$k] = isset($v['rule']) ? $v['rule'] : '';
        }
        Validator::check($data, $rules);
        // option: default => null, ignore => [], json => false, settype = ''
        foreach ($fieldList as $k => $v) {
            if (isset($data[$k])) {
                $val = $data[$k];
                if (!empty($v['ignore']) && in_array($val, $v['ignore'])) {
                    continue;
                }
                if (!empty($v['settype'])) {
                    \settype($val, $v['settype']);
                }
                if (!empty($v['json']) && $v['json']) {
                    $val = \json_decode($val, true);
                }
                $safeData[$k] = $val;
            } else {
                isset($v['default']) && $safeData[$k] = $v['default'];
            }
        }
        return $safeData;
    }

}
