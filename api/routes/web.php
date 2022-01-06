<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
 */

$router->post('/login', ['uses' => 'AuthController@login']);

$router->group(
    ['middleware' => ['jwt', 'permission']],
    function ($router) {

        // Auth
        $router->get('/renewToken', ['uses' => 'AuthController@renewToken']);
        $router->get('/auth/jwt', ['uses' => 'AuthController@jwtInfo']);
        $router->get('/auth/menu', ['uses' => 'AuthController@menu']);
        $router->post('/auth/password', ['uses' => 'AuthController@changPwd']);
        $router->get('/group', ['uses' => 'AuthController@groupList']);
        $router->post('/group', ['uses' => 'AuthController@addGroup']);
        $router->post('/group/edit', ['uses' => 'AuthController@editGroup']);
        $router->post('/group/user', ['uses' => 'AuthController@setGroupUser']);
        $router->post('/group/menu', ['uses' => 'AuthController@setGroupMenu']);
        $router->get('/menu', ['uses' => 'AuthController@menuTree']);
        $router->get('/menu/info', ['uses' => 'AuthController@menuInfo']);
        $router->post('/menu/group', ['uses' => 'AuthController@setMenuGroup']);

        // User
        $router->get('/user/self', ['uses' => 'UserController@self']);
        $router->post('/user/self', ['uses' => 'UserController@editSelf']);
        $router->get('/user/list', ['uses' => 'UserController@list']);
        $router->get('/user/info', ['uses' => 'UserController@info']);
        $router->post('/user/add', ['uses' => 'UserController@add']);
        $router->post('/user/update', ['uses' => 'UserController@updateInfo']);
        $router->post('/user/status', ['uses' => 'UserController@status']);
        $router->post('/user/group', ['uses' => 'UserController@setGroup']);

    }
);
