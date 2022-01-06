#!/bin/bash

cd /work

if [ ! -d './dist' ]; then
    mkdir dist
fi

if [ ! -d './node_modules' ]; then
    npm install
fi

npm run dev