<?php

namespace App\Http\Logic;

use App\Exceptions\LogicException;

class Logic
{
    private static $_instances = [];

    private function __construct()
    {}

    /**
     * Singleton Pattern
     * @return static
     */
    public static function instance()
    {
        $class = get_called_class();
        if (!isset(self::$_instances[$class])) {
            self::$_instances[$class] = new static();
        }
        return self::$_instances[$class];
    }

    protected function throw(int $code = -1, string $msg = '', array $data = [], array $params = [])
    {
        throw new LogicException($code, $msg, $data, $params);
    }
}
