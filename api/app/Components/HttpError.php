<?php

namespace App\Components;

class HttpError extends \Exception

{
    protected $_errorData = [];

    const MSG_MAP = [
        // app
        -1 => 'unkown',
        0 => '',
        1 => 'Parameter error: %s',
        2 => 'Mysql error!',

        // auth
        10001 => 'JWT Token Not Found!',
        10002 => 'Invalid JWT Token',
        10003 => 'Account not found / Password error!',
        10004 => 'Old password error! / password change error!',
        10005 => 'Account is disabled ! please contact administrator !',
        10006 => 'Add account error! ',
        10007 => 'Account user exist!',
        10008 => 'Account password wrong! will ban with %s chance left !',
        10009 => 'Account password retry too much ! set disabled !',

    ];

    const MSG_MAP_CN = [];

    public function __construct(int $code, string $msg = '', array $data = [], array $params = [], string $lang = '')
    {
        $this->_errorData = $data;
        !$msg && $msg = self::getErrorMsg($code, $params, $lang);
        return parent::__construct($msg, $code);
    }

    public static function getErrorMsg(int $code = 0, array $params = [], string $lang = ''): string
    {
        $msgMap = self::getMap($lang);
        $msg = empty($msgMap[$code]) ? 'Unknow Error' : $msgMap[$code];
        return $params ? vsprintf($msg, $params) : $msg;
    }

    public static function getMap(string $lang = ''): array
    {
        $map = $lang != '' ? \constant('self::MSG_MAP_' . strtoupper($lang)) : self::MSG_MAP;
        return empty($map) ? [] : $map;
    }

    public function getErrorData(): array
    {
        return $this->_errorData;
    }

}
