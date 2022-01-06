<?php

namespace App\Components;

use App\Components\Component;
use Illuminate\Support\Facades\Validator as ValidatorFacades;
use Illuminate\Validation\ValidationException;

class Validator extends Component
{
    const MSG_TEMPLATE = [
        'required' => 'The field [:attribute] is required.',
        'size' => 'The field [:attribute] must be exactly [:size].',
        'between' => 'The field [:attribute] value :input is not between [:min - :max].',
        'in' => 'The field [:attribute] must be one of the following types: [:values]',
        'json' => 'The field [:attribute] must be json format',
        'array' => 'The field [:attribute] must be array or object',
        'integer' => 'The field [:attribute] must be integer',
        'string' => 'The field [:attribute] must be string',
        'max' => 'The field [:attribute] must <= [:max]',
        'min' => 'The field [:attribute] must >= [:min]',
    ];

    public static function check(array $data = [], array $rules = [], bool $throw = true): array
    {
        $validator = ValidatorFacades::make($data, $rules, self::MSG_TEMPLATE);
        if ($validator->fails()) {
            if ($throw) {
                throw new ValidationException($validator);
            }
            return $validator->errors()->all();
        }
        return [];
    }
}
