#!/bin/bash

docker run \
    --rm \
    --name admin-web-build \
    -v $(pwd):/work \
    node:14-slim \
    bash -c "chmod +x /work/docker/build.sh && /work/docker/build.sh"
