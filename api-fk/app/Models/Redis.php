<?php

namespace App\Models;

class Redis
{
    private static $_instances = [];
    private static $_connections = [];

    const CONN_DEFAULT = 'default';
    const CONN_CACHE = 'cache';

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

    /**
     * get redis object(phpredis)
     * @param string $key
     * @return \Redis|null
     */
    public function connect(string $key = 'default')
    {
        if (isset(self::$_connections[$key])) {
            return self::$_connections[$key];
        }

        $config = config('database.redis');

        if (!isset($config[$key])) {
            return null;
        }

        try {
            $redis = new \Redis();
            $redis->connect($config[$key]['host'], $config[$key]['port']);
            if (!empty($config[$key]['password'])) {
                $redis->auth($config[$key]['password']);
            }
            if (isset($config[$key]['db'])) {
                $redis->select($config[$key]['db']);
            }
            if (isset($config[$key]['prefix'])) {
                $redis->setOption(\Redis::OPT_PREFIX, $config[$key]['prefix']);
            }
            self::$_connections[$key] = $redis;
        } catch (\Throwable $th) {
            // TODO: log connect error
            throw new \Exception('Redis connect error: ' . $th->getMessage());
            return null;
        }

        return self::$_connections[$key];
    }

}
