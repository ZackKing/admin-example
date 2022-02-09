#!/bin/bash

docker run \
    --rm \
    --name admin-web-dev \
    -v $(pwd):/work \
    -p 8081:8081 \
    node:17-slim \
    bash -c "chmod +x /work/docker/dev.sh && /work/docker/dev.sh"
