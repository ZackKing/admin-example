<?php

namespace App\Http\Middleware;

use App\Components\HttpError;
use App\Http\Logic\Auth as AuthLogic;
use Closure;
use Illuminate\Http\Request;

class JWTMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle(Request $request, Closure $next)
    {
        $token = $request->headers->get('ADMIN-TOKEN');
        if (\is_null($token)) {
            return $this->NotFoundTokenError();
        }
        $payload = AuthLogic::getInstance()->parseToken($token);
        if (!$payload) {
            return $this->InvalidTokenError();
        }
        $request->jwt = $payload;
        return $next($request);
    }

    protected function NotFoundTokenError()
    {
        return response()->json([
            "code" => 10000,
            "msg" => HttpError::getErrorMsg(10000),
            "data" => [],
        ], 200);
    }

    protected function InvalidTokenError()
    {
        return response()->json([
            "code" => 10000,
            "msg" => HttpError::getErrorMsg(10001),
            "data" => [],
        ], 200);
    }

}
