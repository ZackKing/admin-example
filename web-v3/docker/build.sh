#!/bin/bash

cd /work

if [ ! -d './dist' ]; then
    mkdir dist
fi

if [ ! -d './node_modules' ]; then
    yarn install
fi

yarn build