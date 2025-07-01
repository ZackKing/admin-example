<?php

namespace App\Exceptions;

use \Exception;

class LogicException extends Exception
{
    protected $_data = [];

    const MSG_MAP = [
        // app
        -1 => 'unkown',
        0 => '',
        1 => 'Parameter error: %s',
        2 => 'Mysql error!',

        // auth
        10000 => 'Internal auth error',
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

    public function __construct(int $code, string $msg = '', array $data = [], array $params = [])
    {
        $this->_data = $data;
        !$msg && $msg = self::getMsg($code, $params);
        return parent::__construct($msg, $code);
    }

    public static function throw(int $code, string $msg = '', array $data = [], array $params = [])
    {
        throw new self($code, $msg, $data, $params);
    }

    public static function getMsg(int $code = 0, array $params = []): string
    {
        $msg = empty(self::MSG_MAP[$code]) ? 'Unknow Error' : self::MSG_MAP[$code];
        return $params ? vsprintf($msg, $params) : $msg;
    }

    public function getData(): array
    {
        return $this->_data;
    }

}
