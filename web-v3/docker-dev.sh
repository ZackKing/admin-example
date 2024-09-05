#!/bin/bash

docker run \
    --rm \
    --name admin-web-dev \
    -v $(pwd):/work \
    -p 8088:8088 \
    node:20-slim \
    bash -c "cd /work && yarn && yarn dev"
