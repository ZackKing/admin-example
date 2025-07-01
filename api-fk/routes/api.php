<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\UserController;

Route::middleware(['access'])->group(function () {

    Route::post('/login', [AuthController::class, 'login']);

    Route::middleware(['auth'])->group(function () {

        // Auth
        Route::get('/renewToken', [AuthController::class, 'renewToken']);
        Route::get('/auth/jwt', [AuthController::class, 'jwtInfo']);
        Route::get('/auth/menu', [AuthController::class, 'menu']);
        Route::post('/auth/password', [AuthController::class, 'changPwd']);
        Route::get('/group', [AuthController::class, 'groupList']);
        Route::post('/group', [AuthController::class, 'addGroup']);
        Route::post('/group/edit', [AuthController::class, 'editGroup']);
        Route::post('/group/user', [AuthController::class, 'setGroupUser']);
        Route::post('/group/menu', [AuthController::class, 'setGroupMenu']);
        Route::get('/menu', [AuthController::class, 'menuTree']);
        Route::get('/menu/info', [AuthController::class, 'menuInfo']);
        Route::post('/menu/group', [AuthController::class, 'setMenuGroup']);

        // User
        Route::get('/user/self', [UserController::class, 'self']);
        Route::post('/user/self', [UserController::class, 'editSelf']);
        Route::get('/user/list', [UserController::class, 'list']);
        Route::get('/user/info', [UserController::class, 'info']);
        Route::post('/user/add', [UserController::class, 'add']);
        Route::post('/user/update', [UserController::class, 'updateInfo']);
        Route::post('/user/status', [UserController::class, 'status']);
        Route::post('/user/group', [UserController::class, 'setGroup']);

    });

});
