<?php

namespace App\Http\Middleware;

use App\Components\JWTHelper;
use App\Exceptions\LogicException;
use Closure;
use Illuminate\Http\Request;

class Auth
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $r
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle(Request $r, Closure $next)
    {
        $token = $r->headers->get('ADMIN-TOKEN');
        if (\is_null($token)) {
            LogicException::throw(10000);
        }
        $payload = JWTHelper::parseToken($token);
        if (!$payload || empty($payload['uid'])) {
            LogicException::throw(10001);
        }
        $r->offsetSet('sys_auth', $payload);
        $r->offsetSet('sys_uid', $payload['uid']);
        return $next($r);
    }

}
