<?php
namespace App\Exceptions;

use App\Exceptions\LogicException;
use Illuminate\Auth\Access\AuthorizationException;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Validation\ValidationException;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Throwable;

class Handler extends ExceptionHandler
{
    protected static $_dontReport = [
        AuthorizationException::class,
        HttpException::class,
        ModelNotFoundException::class,
        ValidationException::class,
        NotFoundHttpException::class,
        MethodNotAllowedHttpException::class,
        LogicException::class,
    ];

    static function handle(Throwable $e): \Illuminate\Http\JsonResponse
    {
        self::_report2external($e);

        $res = [
            'code' => -1,
            'msg'  => 'server error',
            'ts'   => time(),
        ];

        if (config('app.debug')) {
            $res['msg']       = $e->getMessage();
            $res['exception'] = get_class($e);
            $res['trace']     = array_slice(explode('#', $e->getTraceAsString()), 1, 5);
        }

        if ($e instanceof LogicException) {
            $res['code'] = $e->getCode();
            $res['msg']  = $e->getMessage();
            $res['data'] = $e->getData();
            return response()->json($res, 200);
        } else if ($e instanceof NotFoundHttpException || $e instanceof MethodNotAllowedHttpException) {
            $res['code'] = 404;
            $res['msg']  = 'Api not found';
            return response()->json($res, 404);
        } else if ($e instanceof ValidationException) {
            $res['code'] = 1;
            $res['msg']  = $e->validator->errors()->all(':key => :message');
            return response()->json($res, 200);
        }

        return response()->json($res, 500);
    }

    static function _report2external(Throwable $e): void
    {
        foreach (self::$_dontReport as $type) {
            if ($e instanceof $type) {
                return;
            }
        }

        // TODO: notice to external message service
    }


}
