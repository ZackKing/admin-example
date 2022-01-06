<?php

namespace App\Model;

class Redis
{
    private static $_instances = [];
    private static $_connections = [];

    const CONN_DEFAULT = 'default';

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

    /**
     * get redis object(phpredis)
     * @param string $key
     * @return \Redis|null
     */
    public function getConnection(string $key = 'default')
    {
        if (isset(self::$_connections[$key])) {
            return self::$_connections[$key];
        }

        $config = config('redis');
        if (!isset($config[$key])) {
            return null;
        }

        try {
            $redis = new \Redis();
            $redis->connect($config[$key]['host'], $config[$key]['port']);
            if (isset($config[$key]['password'])) {
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

    public function defaultCache(string $key, $data = null, int $ex = 3600)
    {
        return $this->cache(self::CONN_DEFAULT, $key, $data, $ex);
    }

    public function cache(string $conn, string $key, $data = null, int $ex = 3600)
    {
        if (!$key || !$conn) {
            return false;
        }
        $r = $this->getConnection($conn);
        if (is_null($r)) {
            return false;
        }

        if (\is_null($data)) { // get
            $data = $r->get($key);
            return $data === false ? false : \json_decode($data, true);
        } else { // set
            return $r->setex($key, $ex, \json_encode($data));
        }
    }

}
