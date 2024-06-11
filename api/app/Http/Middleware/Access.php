<?php

namespace App\Http\Middleware;

use App\Components\Helper;
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
        // TOOD: check api auth
        $r->offsetSet('sys_access_id', date('YmdHis', time()) . uniqid());
        $this->_log($r);
        return $next($r);
    }

    private function _log(Request $r)
    {
        if (!env('APP_DEBUG', false) && $r->method() == 'GET') { // don't log GET
            return;
        }

        $log = [
            'access_id' => $r->offsetGet('sys_access_id'),
            'method' => $r->method(),
            'path' => $r->path(),
            'header' => json_encode($r->header()),
            'query' => json_encode($r->query->all()),
            'body' => json_encode($r->json()->all() ?: $r->input()),
            'ip' => Helper::realIp(),
        ];
        // TODO: 投递MQ
        Log::info('access_log', $log);
    }

}
