<?php

namespace App\Components;

class Component
{
    private static $_instances = [];

    /**
     * Singleton
     * @return static
     */
    public static function getInstance()
    {
        $class = get_called_class();
        if (!isset(self::$_instances[$class])) {
            self::$_instances[$class] = new static();
        }
        return self::$_instances[$class];
    }

    protected function __construct()
    {}

}
