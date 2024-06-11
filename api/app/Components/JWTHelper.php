<?php

namespace App\Components;

use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Support\Facades\Log;

class JWTHelper extends Component
{
    public static function createToken(int $uid): string
    {
        if ($uid < 1) {
            return '';
        }
        $conf = config('app.jwt');
        $signTime = time();
        $token = JWT::encode([
            'uid' => $uid,
            'iss' => $conf['iss'],
            'iat' => $signTime,
            'exp' => $signTime + $conf['exp'],
        ],
            $conf['secret'],
            $conf['algorithm']
        );
        return $token;
    }

    public static function parseToken(string $token): array | false
    {
        $conf = config('app.jwt');
        try {
            $decoded = JWT::decode($token, new Key($conf['secret'], $conf['algorithm']));
            return (array) $decoded;
        } catch (\Throwable $th) {
            Log::error('jwt_parse_error', ['msg' => $th->getMessage()]);
            return false;
        }
    }

}
