<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Log;

class TestConsole extends Command
{
    protected $signature = 'cmd:test';
    protected $description = 'for test code';

    public function __construct()
    {
        parent::__construct();
    }

    public function handle()
    {
        try {
            $this->test();
        } catch (\Throwable$th) {
            Log::error($th->getMessage());
        }
    }

    public function test()
    {
        echo 'Hello Admin';
    }
}
