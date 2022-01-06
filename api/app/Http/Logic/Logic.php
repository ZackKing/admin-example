<?php

namespace App\Http\Logic;

use App\Components\HttpError;

class Logic
{
    private static $_instances = [];

    private function __construct()
    {}

    /**
     * Singleton Pattern
     * @return static
     */
    public static function getInstance()
    {
        $class = get_called_class();
        if (!isset(self::$_instances[$class])) {
            self::$_instances[$class] = new static();
        }
        return self::$_instances[$class];
    }

    protected function throwHttpError(int $code = -1, string $msg = '', array $data = [], array $params = [])
    {
        throw new HttpError($code, $msg, $data, $params);
    }
}
