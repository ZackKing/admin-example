<?php

namespace App\Console;

use Illuminate\Console\Scheduling\Schedule;
use Laravel\Lumen\Console\Kernel as ConsoleKernel;

class Kernel extends ConsoleKernel
{
    protected $commands = [
        'App\Console\Commands\TestConsole',
    ];

    protected function schedule(Schedule $schedule)
    {
        // $schedule->command('cmd:test')->cron('*/30 * * * *');
    }
}
