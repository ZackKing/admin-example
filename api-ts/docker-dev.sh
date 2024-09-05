#!/bin/bash

if [ ! -d './node_modules' ];then
  runShell='cd /work && yarn && yarn dev-w'
else
  runShell='cd /work && yarn dev-w'
fi

docker run \
  --rm \
  --name admin-service \
  -v $(pwd):/work \
  -p 127.0.0.1:3366:3366 \
  -d \
  node:20-slim \
  bash -c "$runShell"
