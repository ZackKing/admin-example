#!/bin/bash

# for watch
if [ ! -d node_modules ]; then
    npm install
    if [ $? -ne 0 ]; then
        echo "❌ npm install error"
        exit 1
    fi
fi

php artisan octane:frankenphp \
    --host=0.0.0.0 \
    --port=3088 \
    --workers=4 \
    --max-requests=100\
    --watch
