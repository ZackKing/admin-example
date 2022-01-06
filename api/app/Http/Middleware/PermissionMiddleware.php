<?php

namespace App\Http\Middleware;

use App\Components\Helper;
use App\Model\AccessLog;
use Closure;
use Illuminate\Http\Request;

class PermissionMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle(Request $r, Closure $next)
    {
        // TOOD: check api auth

        $r->realIp = Helper::realIp();
        $r->accessLogId = $this->_log($r);
        return $next($r);
    }

    private function _log(Request $r): int
    {
        // don't log GET
        if ($r->method() == 'GET') {
            return 0;
        }
        $log = [
            'uid' => $r->jwt['uid'],
            'method' => $r->method(),
            'path' => $r->path(),
            'header' => \json_encode($r->header()),
            'query' => \json_encode($r->query->all()),
            'body' => \json_encode($r->json()->all()),
            'ip' => $r->realIp,
        ];
        return AccessLog::getInstance()->insert($log);
    }

}
