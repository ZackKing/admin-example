#!/bin/bash

docker run \
    --rm \
    --name admin-web-build \
    -v $(pwd):/work \
    node:18-slim \
    bash -c "/work/docker/build.sh"
