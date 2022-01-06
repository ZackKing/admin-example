<?php

namespace App\Exceptions;

use App\Components\HttpError as HttpLogicError;
use Illuminate\Auth\Access\AuthorizationException;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Validation\ValidationException;
use Laravel\Lumen\Exceptions\Handler as ExceptionHandler;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;

class Handler extends ExceptionHandler
{
    /**
     * A list of the exception types that should not be reported.
     *
     * @var array
     */
    protected $dontReport = [
        AuthorizationException::class,
        HttpException::class,
        ModelNotFoundException::class,
        ValidationException::class,
        NotFoundHttpException::class,
        MethodNotAllowedHttpException::class,
        HttpLogicError::class,
    ];

    /**
     * Report or log an exception.
     *
     * This is a great spot to send exceptions to Sentry, Bugsnag, etc.
     *
     * @param  \Throwable  $exception
     * @return void
     *
     * @throws \Exception
     */
    public function report(\Throwable$e)
    {
        parent::report($e);
        if ($this->shouldReport($e)) {
            // report for some notify server
        }
    }

    /**
     * Render an exception into an HTTP response.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Throwable  $exception
     * @return \Illuminate\Http\Response|\Illuminate\Http\JsonResponse
     *
     * @throws \Throwable
     */
    public function render($request, \Throwable $th)
    {
        $res = [
            'code' => -1,
            'msg' => 'server error',
        ];
        if (config('app.debug')) {
            $res['msg'] = $th->getMessage();
            $res['exception'] = get_class($th);
            $res['trace'] = array_slice($th->getTrace(), 0, 5);
        }

        if ($th instanceof HttpLogicError) { // handle http logic error
            $res['code'] = $th->getCode();
            $res['msg'] = $th->getMessage();
            $res['data'] = $th->getErrorData();
            return $this->_response()->json($res, 200);
        } else if ($th instanceof NotFoundHttpException || $th instanceof MethodNotAllowedHttpException) { // handle 404
            $res['code'] = 404;
            $res['msg'] = 'Api not found';
            return $this->_response()->json($res, 404);
        } else if ($th instanceof ValidationException) { //handle data params error
            $res['code'] = 1;
            $res['msg'] = $th->validator->errors()->all(':key => :message');
            return $this->_response()->json($res, 200);
        }
        return $this->_response()->json($res, 200, [], 0);
    }

    private function _response($content = '', $status = 200, array $headers = [])
    {
        $factory = new \Laravel\Lumen\Http\ResponseFactory;
        if (func_num_args() === 0) {
            return $factory;
        }
        return $factory->make($content, $status, $headers);
    }
}
