#!/bin/bash

docker run \
    --rm \
    -v $(pwd):/work \
    zackzaaack/apidoc -o /work/public/apidoc  -i /work/app/Http/Controllers