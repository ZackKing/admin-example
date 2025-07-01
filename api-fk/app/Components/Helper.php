<?php

namespace App\Components;

class Helper extends Component
{
    /**
     * convert object to array 
     * @param object $object
     * @return array
     */
    public static function obj2arr(object $object): array
    {
        if (is_object($object)) {
            foreach ($object as $key => $value) {
                $array[$key] = $value;
            }
        } else {
            $array = $object;
        }
        return $array;
    }

    /**
     * field filter
     * @param array $data
     * @param array $fields
     * @return array
     */
    public static function fieldFilter(array $data = [], array $fields = []): array
    {
        $safeData = [];
        foreach ($fields as $field) {
            if (isset($data[$field])) {
                $safeData[$field] = $data[$field];
            }
        }
        return $safeData;
    }

    /**
     * change value be key
     * @param array $list
     * @param string $key
     * @return array
     */
    public static function value2key(array $list = [], string $key = ''): array
    {
        return array_column($list, null, $key);
    }

    /**
     * convert to tree data
     * @param array $items
     * @param string $key
     * @param string $pkey
     * @param string $skey
     * @return void
     */
    public static function genTree(array $list = [], string $key = 'id', string $pkey = 'pid', string $skey = 'sub'): array
    {
        $tree = [];
        $items = self::value2key($list, $key);
        foreach ($items as $item) {
            if (isset($items[$item[$pkey]])) {
                $items[$item[$pkey]][$skey][] = &$items[$item[$key]];
            } else {
                $tree[] = &$items[$item[$key]];
            }
        }
        return $tree;
    }

    /**
     * convert assoc to array
     * @param array $hash
     * @return array
     */
    public static function hash2list(array $assoc): array
    {
        $list = [];
        foreach ($assoc as $v) {
            $list[] = $v;
        }
        return $list;
    }

    /**
     * gat ervery day date between start time and end time
     * @param  int  $stime  start time
     * @param  int  $etime  end time
     * @return Array
     */
    public static function getDailyDate(int $stime, int $etime, string $format = 'Ymd', bool $desc = false): array
    {
        $days = ($etime - $stime) / 86400;
        $dates = [];
        for ($i = 0; $i <= $days; $i++) {
            $dates[] = date($format, $stime + (86400 * $i));
        }
        return $desc ? array_reverse($dates) : $dates;
    }

    /**
     * gen md5 16 string
     * @param string $str
     * @return string
     */
    public static function md5_16(string $str): string
    {
        return substr(md5($str), 8, 16);
    }

    public static function listValueMerge(array $list, string $field, bool $toList = true): array
    {
        $res = [];
        foreach ($list as $val) {
            empty($res[$val[$field]]) && $res[$val[$field]] = [$field => $val[$field]];
            foreach ($val as $k => $v) {
                if ($k != $field) {
                    $res[$val[$field]][$k] = $v;
                }
            }
        }
        return $toList ? self::hash2list($res) : $res;
    }

    public static function randStr(int $length = 6, string $str = '')
    {
        if (!$str) {
            $str = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
        }
        $s = '';
        for ($i = 0; $i < $length; $i++) {
            $s .= $str[mt_rand(0, strlen($str) - 1)];
        }
        return $s;
    }

    public static function realIp(): string
    {
        $realip = '';
        if (isset($_SERVER)) {
            if (isset($_SERVER['HTTP_X_FORWARDED_FOR'])) {
                $arr = explode(',', $_SERVER['HTTP_X_FORWARDED_FOR']);
                foreach ($arr as $ip) {
                    $ip = trim($ip);
                    if ($ip != 'unknown') {
                        $realip = $ip;
                        break;
                    }
                }
            } elseif (isset($_SERVER['HTTP_CLIENT_IP'])) {
                $realip = $_SERVER['HTTP_CLIENT_IP'];
            } else {
                if (isset($_SERVER['REMOTE_ADDR'])) {
                    $realip = $_SERVER['REMOTE_ADDR'];
                } else {
                    $realip = '0.0.0.0';
                }
            }
        } else {
            if (getenv('HTTP_X_FORWARDED_FOR')) {
                $realip = getenv('HTTP_X_FORWARDED_FOR');
            } elseif (getenv('HTTP_CLIENT_IP')) {
                $realip = getenv('HTTP_CLIENT_IP');
            } else {
                $realip = getenv('REMOTE_ADDR');
            }
        }
        preg_match("/[\d\.]{7,15}/", $realip, $onlineip);
        $realip = !empty($onlineip[0]) ? $onlineip[0] : '0.0.0.0';
        return $realip;
    }

    public static function objList2arr($list): array
    {
        return json_decode(json_encode($list), true);
    }

}
