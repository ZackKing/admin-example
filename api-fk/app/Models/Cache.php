<?php

namespace App\Models;

class Cache
{
    public static function get(string $k, $default = false)
    {
        $k = self::_genKey($k);
        if(!$r = Redis::instance()->connect(Redis::CONN_CACHE)) {
            return $default;
        }
        $data = $r->get($k);
        return $data === false ? $default : json_decode($data, true);
    }

    public static function set(string $k, $data, int $ex = 300)
    {
        $k = self::_genKey($k);
        if(!$r = Redis::instance()->connect(Redis::CONN_CACHE)) {
            return false;
        }
        return $r->setex($k, $ex, json_encode($data, JSON_UNESCAPED_UNICODE));
    }

    private static function _genKey(string $k)
    {
        return 'cache:' . $k;
    }

}
