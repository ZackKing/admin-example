#!/bin/bash

docker run \
    --rm \
    --name admin-web-build \
    -v $(pwd):/work \
    node:20-slim \
    bash -c "cd /work && yarn && yarn build"
