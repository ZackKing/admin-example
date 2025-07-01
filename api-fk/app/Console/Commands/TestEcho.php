<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class TestEcho extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'test:echo';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Test echo stript';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        echo 'hello admin';
    }
}
