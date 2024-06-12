<?php

namespace App\Http\Middleware;

use App\Components\Helper;
use App\Models\Admin\AccessLog;
use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Log;

class Access
{
    /**
     * Handle an incoming request.
     *
     * @param \Illuminate\Http\Request $request
     * @param \Closure $next
     * @return mixed
     */
    public function handle(Request $r, Closure $next)
    {
        $r->offsetSet('sys_access_id', uniqid(date('YmdH', time())));
        $res = $next($r);
        $this->_log($r, $res->original ?? []);
        return $res;
    }

    private function _log(Request $r, array $resData = [])
    {
        if (!env('APP_DEBUG', false) && $r->method() == 'GET') { // don't log GET
            return;
        }

        $log = [
            'access_id' => $r->offsetGet('sys_access_id', ''),
            'uid' => $r->offsetGet('sys_uid', 0),
            'method' => $r->method(),
            'path' => $r->path(),
            'header' => json_encode($r->header()),
            'query' => json_encode($r->query->all(), JSON_UNESCAPED_UNICODE),
            'body' => json_encode($r->json()->all() ?: $r->input(), JSON_UNESCAPED_UNICODE),
            'ip' => Helper::realIp(),
            'response' => json_encode($resData, JSON_UNESCAPED_UNICODE),
        ];
        AccessLog::instance()->insert($log);
        Log::info('access_log', $log);
    }

}
