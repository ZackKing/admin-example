#!/bin/bash

docker run \
    --rm \
    --name admin-web-dev \
    -v $(pwd):/work \
    -p 8080:8080 \
    node:14-slim \
    bash -c "chmod +x /work/docker/dev.sh && /work/docker/dev.sh"
